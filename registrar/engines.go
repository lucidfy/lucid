package registrar

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/lucidfy/lucid/pkg/facade/routes"
)

var Engines = map[string]func() http.Handler{
	"mux":   mux,
	"fiber": fiber,
}

func mux() http.Handler {
	return routes.NetHttp().Register(&Routes)
}

func fiber() http.Handler {
	return adaptor.FiberApp(routes.Fiber(&Routes).App)
}
