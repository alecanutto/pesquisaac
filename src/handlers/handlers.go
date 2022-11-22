package handlers

import (
	"github.com/alecanutto/pesquisaac/src/handlers/health"

	"github.com/OlaIsaac/horcrux/modules/server"
	"github.com/alecanutto/pesquisaac/src/domain"
)

type HandlerContainer struct {
	Handlers []server.Handler
	Routers  []server.Router
}

func NewHandlerContainer() HandlerContainer {
	return HandlerContainer{
		Handlers: []server.Handler{
			health.NewHealthHandler(),
		},
		Routers: []server.Router{
			{
				Key:       domain.API_PUBLIC_KEY,
				Path:      "/",
				IsEnabled: true,
			},
		},
	}
}
