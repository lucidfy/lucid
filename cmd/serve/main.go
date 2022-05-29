package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/lucidfy/lucid/app"
	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/internal/kernel"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/registrar"
	"github.com/lucidfy/lucid/resources/translations"
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
	trans := lang.Load(translations.Languages)

	return map[string]func() http.Handler{
		"mux": func() http.Handler {

			nethttp := routes.NetHttp(trans).
				AddGlobalMiddlewares(app.GlobalMiddleware).
				AddRouteMiddlewares(app.RouteMiddleware)

			nethttp.HttpErrorHandler = handlers.HttpErrorHandler

			nethttp.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				engine := *engines.NetHttp(w, r, trans)
				handlers.HttpErrorHandler(engine, &errors.AppError{
					Code:    http.StatusNotFound,
					Message: "Page not found",
					Error:   fmt.Errorf("404 page not found"),
				}, nil)
			})

			nethttp.Router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				engine := *engines.NetHttp(w, r, trans)
				handlers.HttpErrorHandler(engine, &errors.AppError{
					Code:    http.StatusMethodNotAllowed,
					Message: "Method not allowed",
					Error:   fmt.Errorf("405 method not allowed"),
				}, nil)
			})

			return nethttp.Register(&registrar.Routes)
		},

		"fiber": func() http.Handler {
			fiber := routes.Fiber(trans).
				AddGlobalMiddlewares(app.GlobalMiddleware).
				AddRouteMiddlewares(app.RouteMiddleware)

			fiber.HttpErrorHandler = handlers.HttpErrorHandler

			return adaptor.FiberApp(fiber.Register(&registrar.Routes).App)
		},
	}
}

func printEnvDefaults() {
	log.Println("Defaults: ")
	log.Println(fmt.Sprintf(` -> Lucid Root: %s`, path.Load().BasePath("")))
	log.Println(fmt.Sprintf(` -> Environment: %s`, os.Getenv("APP_ENV")))
	log.Println(fmt.Sprintf(` -> Timezone: %s`, os.Getenv("APP_TIMEZONE")))
	log.Println(fmt.Sprintf(` -> Language: %s`, os.Getenv("APP_LANGUAGE")))
	log.Println(fmt.Sprintf(` -> Scheme: %s`, os.Getenv("SCHEME")))
	log.Println(fmt.Sprintf(` -> Host: %s`, os.Getenv("HOST")))
	log.Println(fmt.Sprintf(` -> Port: %s`, os.Getenv("PORT")))
}

func main() {
	printEnvDefaults()

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
