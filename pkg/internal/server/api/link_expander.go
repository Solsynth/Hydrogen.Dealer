package api

import (
	"encoding/base64"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"github.com/gofiber/fiber/v2"
	"sync"
)

var inProgress sync.Map

func getLinkMeta(c *fiber.Ctx) error {
	targetEncoded := c.Params("target")
	targetRaw, _ := base64.URLEncoding.DecodeString(targetEncoded)

	if ch, loaded := inProgress.LoadOrStore(targetEncoded, make(chan struct{})); loaded {
		// If the request is already in progress, wait for it to complete
		<-ch.(chan struct{})
	} else {
		// If this is the first request, process it and signal others
		defer func() {
			close(ch.(chan struct{}))
			inProgress.Delete(targetEncoded)
		}()
	}

	if meta, err := services.LinkExpand(string(targetRaw)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else {
		return c.JSON(meta)
	}
}
