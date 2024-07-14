package hyper

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"google.golang.org/grpc"
)

func (v *HyperConn) DoAuthenticate(atk, rtk string) (acc *proto.UserInfo, accessTk string, refreshTk string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var in *grpc.ClientConn
	in, err = v.GetServiceGrpcConn(ServiceTypeAuthProvider)
	if err != nil {
		return
	}

	var reply *proto.AuthReply
	reply, err = proto.NewAuthClient(in).Authenticate(ctx, &proto.AuthRequest{
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

func (v *HyperConn) CheckPermGranted(atk string, key string, val []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	in, err := v.GetServiceGrpcConn(ServiceTypeAuthProvider)
	if err != nil {
		return err
	}

	reply, err := proto.NewAuthClient(in).EnsurePermGranted(ctx, &proto.CheckPermRequest{
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
