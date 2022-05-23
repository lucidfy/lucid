package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/lucidfy/lucid/internal/kernel"
	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/registrar"
)

func init() {
	env.LoadEnv()

	var scheme string
	var host string
	var port string
	var router_engine string
	flag.StringVar(&scheme, "scheme", os.Getenv("SCHEME"), "Scheme to use")
	flag.StringVar(&host, "host", os.Getenv("HOST"), "Host to use")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "Port to use")
	flag.StringVar(&router_engine, "router-engine", "mux", `By default we are using "mux"`)
	flag.Parse()

	flags := map[string]string{
		"SCHEME":              scheme,
		"HOST":                host,
		"PORT":                port,
		"LUCID_ROUTER_ENGINE": router_engine,
	}

	for k, v := range flags {
		if len(v) != 0 {
			os.Setenv(k, v)
		}
	}
}

func defaultRouterEngines() map[string]func() http.Handler {
	return map[string]func() http.Handler{
		"mux": func() http.Handler {
			return routes.NetHttp().Register(&registrar.Routes)
		},
		"fiber": func() http.Handler {
			return adaptor.FiberApp(routes.Fiber(&registrar.Routes).App)
		},
	}
}

func main() {
	lre := os.Getenv("LUCID_ROUTER_ENGINE")
	handler, ok := defaultRouterEngines()[lre]

	if !ok {
		panic(fmt.Errorf(`[%s] router engine does not exists`, lre))
	}

	kernel.
		New(handler()).
		Run().
		WithGracefulShutdown()
}
