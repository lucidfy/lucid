package routes

import (
	"net/http"

	"github.com/daison12006013/gorvel/app"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/gorilla/mux"
)

type MuxRoutes struct {
	Router   *mux.Router
	Routings *[]Routing
}

func Mux() MuxRoutes {
	mr := MuxRoutes{}
	mr.Router = mux.NewRouter().StrictSlash(true)
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
		if len(route.Resources) > 0 {
			routings = append(routings, resources(route)...)
		}

		if route.Handler != nil {
			routings = append(routings, route)
		}
	}
	return &routings
}

func (sub MuxRoutes) register(route Routing) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		engine := engines.Mux(w, r)
		route.Handler(engine)
	}

	sub.Router.HandleFunc(route.Path, handler).
		Methods(getMethods(route.Method)...).
		Queries(route.Queries...).
		Name(route.Name)

	for _, v := range route.Middlewares {
		sub.routeUse(app.RouteMiddleware[v])
	}
}

func (mr MuxRoutes) routeUse(middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		mr.Router.Use(middleware)
	}
}
