package routes

import (
	"net/http"

	"github.com/daison12006013/gorvel/app"
	"github.com/gorilla/mux"
)

type Middlewares []string
type Queries []string
type Method []string
type Handler func(http.ResponseWriter, *http.Request)
type Resources map[string]Handler
type Routing struct {
	Name        string
	Path        string
	Method      Method
	Queries     Queries
	Handler     Handler
	Resources   map[string]Handler
	Middlewares Middlewares
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func Register() *mux.Router {
	registrar := mux.NewRouter().StrictSlash(true)

	// Register the global middlewares
	routeUse(registrar, app.Middleware...)

	for _, routing := range *Explain() {
		subrouter := registrar.NewRoute().Subrouter()
		register(subrouter, routing)
	}

	return registrar
}

func Explain() *[]Routing {
	routings := []Routing{}
	for _, route := range *Routes() {
		if len(route.Resources) > 0 {
			routings = append(routings, resources(route)...)
		}

		if route.Handler != nil {
			routings = append(routings, route)
		}
	}
	return &routings
}

func resources(route Routing) []Routing {
	routings := []Routing{}
	for action, handler := range route.Resources {
		switch action {
		case "index":
			routings = append(routings, Routing{
				Path:        route.Path,
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".lists",
			})
		case "create":
			routings = append(routings, Routing{
				Path:        route.Path + "/create",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".create",
			})
		case "store":
			routings = append(routings, Routing{
				Path:        route.Path,
				Handler:     handler,
				Method:      Method{"POST"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".store",
			})
		case "show":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".show",
			})
		case "edit":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}/edit",
				Handler:     handler,
				Method:      Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".edit",
			})
		case "update":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}",
				Handler:     handler,
				Method:      Method{"PUT"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".update",
			})
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}/update",
				Handler:     handler,
				Method:      Method{"POST"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".update.via-post",
			})
		case "destroy":
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}",
				Handler:     handler,
				Method:      Method{"DELETE"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".destroy",
			})
			routings = append(routings, Routing{
				Path:        route.Path + "/{id}/delete",
				Handler:     handler,
				Method:      Method{"POST"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".destroy.via-post",
			})
		}
	}
	return routings
}

func register(
	subrouter *mux.Router,
	route Routing,
) {
	subrouter.HandleFunc(route.Path, route.Handler).
		Methods(getMethods(route.Method)...).
		Queries(route.Queries...).
		Name(route.Name)

	for _, v := range route.Middlewares {
		routeUse(subrouter, app.RouteMiddleware[v])
	}
}

func routeUse(route *mux.Router, middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		route.Use(middleware)
	}
}

func getMethods(methods Method) Method {
	if len(methods) == 0 {
		methods = Method{http.MethodGet}
	}
	return methods
}
