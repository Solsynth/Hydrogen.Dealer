package api

import (
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/models"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog/log"
)

func listenWebsocket(c *websocket.Conn) {
	user := c.Locals("user").(models.Account)

	// Push connection
	services.ClientRegister(user, c)
	log.Debug().Uint("user", user.ID).Msg("New websocket connection established...")

	// Event loop
	var err error

	for {
		if _, _, err = c.ReadMessage(); err != nil {
			break
		}
	}

	// Pop connection
	services.ClientUnregister(user, c)
	log.Debug().Uint("user", user.ID).Msg("A websocket connection disconnected...")
}
