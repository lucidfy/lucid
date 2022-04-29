package routes

import (
	"net/http"

	"github.com/daison12006013/lucid/app"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/gorilla/mux"
)

type FiberRoutes struct {
	Router   *mux.Router
	Routings *[]Routing
}

func Fiber() FiberRoutes {
	fr := FiberRoutes{}
	fr.Router = mux.NewRouter().StrictSlash(true)
	return fr
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func (fr FiberRoutes) Register(base *[]Routing) interface{} {
	// Register the global middlewares
	fr.routeUse(app.Middleware...)

	// each routing should be interpreted as subrouter
	// the subrouter in mux isolates each path with
	// a way to register a repetitive middlewares
	for _, routing := range *fr.Explain(base) {
		subfr := FiberRoutes{}
		subfr.Router = fr.Router.NewRoute().Subrouter()
		subfr.register(routing)
	}

	return fr.Router
}

func (fr FiberRoutes) Explain(base *[]Routing) *[]Routing {
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

func (sub FiberRoutes) register(route Routing) {
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

func (fr FiberRoutes) routeUse(middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		fr.Router.Use(middleware)
	}
}
