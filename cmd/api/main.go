package main

import (
	"github.com/alecanutto/pesquisaac/src/handlers"

	"github.com/OlaIsaac/horcrux/modules/server"
)

func main() {
	handlerContainer := handlers.NewHandlerContainer()

	server.NewServer(server.ServerConfig{
		Cors: server.CorsConfig{
			IsEnabled:        true,
			AllowOrigins:     "*",
			AllowMethods:     "GET, POST, PATCH",
			AllowCredentials: false,
		},
		Routers:   handlerContainer.Routers,
		Handlers:  handlerContainer.Handlers,
		HasRender: true,
	}).Start()
}
