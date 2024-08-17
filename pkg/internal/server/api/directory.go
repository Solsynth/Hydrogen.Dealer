package api

import (
	"github.com/spf13/viper"
	"strings"

	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

func listExistsService(c *fiber.Ctx) error {
	services := directory.ListServiceInstance()

	return c.JSON(lo.Map(services, func(item *directory.ServiceInstance, index int) map[string]any {
		return map[string]any{
			"id":    item.ID,
			"type":  item.Type,
			"label": item.Label,
		}
	}))
}

func forwardServiceRequest(c *fiber.Ctx) error {
	serviceType := c.Params("service")

	aliasingMap := viper.GetStringMapString("services.alias")
	if val, ok := aliasingMap[serviceType]; ok {
		serviceType = val
	}

	service := directory.GetServiceInstanceByType(serviceType)

	if service == nil || service.HttpAddr == nil {
		return fiber.NewError(fiber.StatusNotFound, "service not found")
	}

	ogUrl := c.Request().URI().String()
	url := c.OriginalURL()
	prevUrl := url
	url = strings.Replace(url, "/srv/"+serviceType, "/api", 1)
	if prevUrl == url {
		url = strings.Replace(url, "/cgi/"+serviceType, "/api", 1)
	}
	url = "http://" + *service.HttpAddr + url

	log.Debug().
		Str("from", ogUrl).
		Str("to", url).
		Str("service", serviceType).
		Str("id", service.ID).
		Msg("Forwarding request for service...")

	return proxy.Do(c, url)
}
