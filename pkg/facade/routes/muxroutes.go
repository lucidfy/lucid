package routes

import (
	"net/http"

	"github.com/daison12006013/lucid/app"
	"github.com/daison12006013/lucid/app/handlers"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/gorilla/mux"
)

type MuxRoutes struct {
	Router   *mux.Router
	Routings *[]Routing
}

func Mux() MuxRoutes {
	mr := MuxRoutes{}
	mr.Router = mux.NewRouter().StrictSlash(true)
	mr.setUpDefaultErrors()

	return mr
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func (mr MuxRoutes) Register(base *[]Routing) interface{} {
	// Register the global middlewares
	mr.routeUse(app.Middleware...)

	// each routing should be interpreted as subrouter
	// the subrouter in mux isolates each path with
	// a way to register a repetitive middlewares
	for _, routing := range *mr.Explain(base).(*[]Routing) {
		submr := MuxRoutes{}
		submr.Router = mr.Router.NewRoute().Subrouter()
		submr.register(routing)
	}

	return mr.Router
}

func (mr MuxRoutes) Explain(base *[]Routing) interface{} {
	routings := []Routing{}
	for _, route := range *base {
		if len(route.Resources) != 0 {
			routings = append(routings, resources(route)...)
		}

		if route.Handler != nil || len(route.Static) != 0 {
			routings = append(routings, route)
		}
	}

	return &routings
}

func (mr MuxRoutes) register(route Routing) {
	// serve static
	if len(route.Static) != 0 {
		serve := http.FileServer(http.Dir(route.Static))
		mr.Router.
			PathPrefix(route.Path).
			Handler(http.StripPrefix(route.Path, serve))
		return
	}

	// serving handler based
	handler := func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.Mux(w, r)
		e := route.Handler(engine)
		if e != nil {
			handlers.HttpErrorHandler(engine, e)
		}
	}

	if route.Prefix {
		mr.Router.PathPrefix(route.Path).HandlerFunc(handler).
			Methods(getMethods(route.Method)...).
			Queries(route.Queries...).
			Name(route.Name)
	} else {
		mr.Router.HandleFunc(route.Path, handler).
			Methods(getMethods(route.Method)...).
			Queries(route.Queries...).
			Name(route.Name)
	}

	for _, v := range route.Middlewares {
		mr.routeUse(app.RouteMiddleware[v])
	}
}

func (mr MuxRoutes) routeUse(middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		mr.Router.Use(middleware)
	}
}

// on this function, we setup the default 404 and 405 error page
// this will go thru under the app/handlers/error.go
func (mr MuxRoutes) setUpDefaultErrors() {
	mr.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.Mux(w, r)
		handlers.PageNotFound(engine)
	})

	mr.Router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.Mux(w, r)
		handlers.MethodNotAllowed(engine)
	})
}
