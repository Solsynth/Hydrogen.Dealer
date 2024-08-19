package api

import (
	"encoding/base64"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"github.com/gofiber/fiber/v2"
)

func getLinkMeta(c *fiber.Ctx) error {
	targetEncoded := c.Params("target")
	targetRaw, _ := base64.StdEncoding.DecodeString(targetEncoded)

	if meta, err := services.LinkExpand(string(targetRaw)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else {
		return c.JSON(meta)
	}
}
