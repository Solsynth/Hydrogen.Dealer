package server

import (
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/server/api"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/server/exts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type HTTPApp struct {
	app *fiber.App
}

func NewServer() *HTTPApp {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnableIPValidation:    true,
		ServerHeader:          "Hydrogen.Dealer",
		AppName:               "Hydrogen.Dealer",
		ProxyHeader:           fiber.HeaderXForwardedFor,
		JSONEncoder:           jsoniter.ConfigCompatibleWithStandardLibrary.Marshal,
		JSONDecoder:           jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
		BodyLimit:             512 * 1024 * 1024 * 1024, // 512 TiB
		EnablePrintRoutes:     viper.GetBool("debug.print_routes"),
	})

	app.Use(idempotency.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))

	app.Use(logger.New(logger.Config{
		Format: "${status} | ${latency} | ${method} ${path}\n",
		Output: log.Logger,
	}))

	app.Use(exts.AuthMiddleware)
	app.Use(exts.LinkAccountMiddleware)

	api.MapAPIs(app)

	return &HTTPApp{app}
}

func (v *HTTPApp) Listen() {
	if err := v.app.Listen(viper.GetString("bind")); err != nil {
		log.Fatal().Err(err).Msg("An error occurred when starting server...")
	}
}
