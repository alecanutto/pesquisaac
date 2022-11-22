package health

import (
	"net/http"

	"github.com/alecanutto/pesquisaac/src/domain"

	"github.com/OlaIsaac/horcrux/modules/server"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (rh HealthHandler) Routes(r server.ServerRouter) {
	r.Router(domain.API_PUBLIC_KEY, func(f fiber.Router) {
		fr := f.Group("/health")
		fr.Get("/ping", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON("pong")
		})
	})
}
