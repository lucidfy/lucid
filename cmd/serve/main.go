package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/internal/kernel"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/loader"
	"github.com/lucidfy/lucid/registrar"
	"github.com/lucidfy/lucid/resources/translations"
)

func init() {
	env.LoadEnv()

	var scheme string
	var host string
	var port string
	var router_engine string
	var show_defaults bool
	flag.StringVar(&scheme, "scheme", os.Getenv("SCHEME"), "Scheme to use")
	flag.StringVar(&host, "host", os.Getenv("HOST"), "Host to use")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "Port to use")
	flag.StringVar(&router_engine, "router-engine", "mux", `By default we are using "mux"`)
	flag.BoolVar(&show_defaults, "show-defaults", false, "Print out the default env")
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

	showDefaults()
}

func main() {
	lre := os.Getenv("LUCID_ROUTER_ENGINE")

	bootstrap(lre)
}

func bootstrap(e string) {
	trans := lang.Load(translations.Languages)

	lists := map[string]func(){
		"mux": func() {
			nethttp := loader.NetHttp(trans).
				AddGlobalMiddlewares(registrar.GlobalMiddleware).
				AddRouteMiddlewares(registrar.RouteMiddleware)

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

			handler := nethttp.Register(&registrar.Routes)

			kernel.
				NetHttp(handler).
				Run().
				WithGracefulShutdown()
		},
	}

	callback, ok := lists[e]

	if !ok {
		panic(fmt.Errorf(`[%s] router engine does not exists`, e))
	}

	callback()
}

func showDefaults() {
	log.Println("Defaults: ")
	log.Println(fmt.Sprintf(` -> Environment: %s`, os.Getenv("APP_ENV")))
	log.Println(fmt.Sprintf(` -> Timezone: %s`, os.Getenv("APP_TIMEZONE")))
	log.Println(fmt.Sprintf(` -> Language: %s`, os.Getenv("APP_LANGUAGE")))
	log.Println(fmt.Sprintf(` -> Scheme: %s`, os.Getenv("SCHEME")))
	log.Println(fmt.Sprintf(` -> Host: %s`, os.Getenv("HOST")))
	log.Println(fmt.Sprintf(` -> Port: %s`, os.Getenv("PORT")))
}
