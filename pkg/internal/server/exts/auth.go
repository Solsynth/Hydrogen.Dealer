package exts

import (
	"context"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

func AuthMiddleware(c *fiber.Ctx) error {
	var atk string
	if cookie := c.Cookies(hyper.CookieAtk); len(cookie) > 0 {
		atk = cookie
	}
	if header := c.Get(fiber.HeaderAuthorization); len(header) > 0 {
		tk := strings.Replace(header, "Bearer", "", 1)
		atk = strings.TrimSpace(tk)
	}
	if tk := c.Query("tk"); len(tk) > 0 {
		atk = strings.TrimSpace(tk)
	}

	c.Locals("p_token", atk)

	rtk := c.Cookies(hyper.CookieRtk)
	if acc, newAtk, newRtk, err := DoAuthenticate(atk, rtk); err == nil {
		if newAtk != atk {
			SetAuthCookies(c, newAtk, newRtk)
		}
		c.Locals("user", acc)
	}

	return c.Next()
}

func DoAuthenticate(atk, rtk string) (acc *proto.UserInfo, accessTk string, refreshTk string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	in := directory.GetServiceInstanceByType(directory.ServiceTypeAuthProvider)
	if in == nil {
		return
	}
	conn, err := in.GetGrpcConn()
	if err != nil {
		return
	}

	var reply *proto.AuthReply
	reply, err = proto.NewAuthClient(conn).Authenticate(ctx, &proto.AuthRequest{
		AccessToken:  atk,
		RefreshToken: &rtk,
	})
	if err != nil {
		return
	}
	if reply != nil {
		acc = reply.GetInfo().GetInfo()
		accessTk = reply.GetInfo().GetNewAccessToken()
		refreshTk = reply.GetInfo().GetNewRefreshToken()
		if !reply.IsValid {
			err = fmt.Errorf("invalid authorization context")
			return
		}
	}

	return
}

func CheckPermGranted(atk string, key string, val []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	in := directory.GetServiceInstanceByType(directory.ServiceTypeAuthProvider)
	if in == nil {
		return fmt.Errorf("no auth provider found")
	}
	conn, err := in.GetGrpcConn()
	if err != nil {
		return err
	}

	reply, err := proto.NewAuthClient(conn).EnsurePermGranted(ctx, &proto.CheckPermRequest{
		Token: atk,
		Key:   key,
		Value: val,
	})
	if err != nil {
		return err
	} else if !reply.GetIsValid() {
		return fmt.Errorf("missing permission: %s", key)
	}

	return nil
}

func EnsureAuthenticated(c *fiber.Ctx) error {
	if _, ok := c.Locals("p_user").(*proto.UserInfo); !ok {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	return nil
}

func EnsureGrantedPerm(c *fiber.Ctx, key string, val any) error {
	if err := EnsureAuthenticated(c); err != nil {
		return err
	}
	encodedVal, _ := jsoniter.Marshal(val)
	if err := CheckPermGranted(c.Locals("p_token").(string), key, encodedVal); err != nil {
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}
	return nil
}
