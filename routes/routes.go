package routes

import (
	"net/http"

	"github.com/daison12006013/gorvel/app"
	"github.com/daison12006013/gorvel/app/handlers"
	"github.com/gorilla/mux"
)

func Routes() *[]route {
	l := &[]route{
		{
			path:    "/",
			method:  []string{"GET"},
			handler: handlers.Home,
		},
		{
			path:        "/users",
			method:      []string{"POST"},
			handler:     handlers.UserCreate,
			middlewares: []string{"auth"},
		},
		{
			path:        "/users",
			method:      []string{"GET"},
			handler:     handlers.UserLists,
			queries:     []string{},
			middlewares: []string{"auth"},
		},
	}

	return l
}

type route struct {
	path        string
	method      []string
	queries     []string
	handler     func(http.ResponseWriter, *http.Request)
	middlewares []string
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func Register() *mux.Router {
	registrar := mux.NewRouter().StrictSlash(true)

	// * Register the global middlewares
	injectMiddleware(registrar, app.Middleware...)

	for _, route := range *Routes() {
		subrouter := registrar.NewRoute().Subrouter()

		subrouter.HandleFunc(route.path, route.handler).
			Methods(getMethods(route.method)...).
			Queries(route.queries...)

		for _, v := range route.middlewares {
			injectMiddleware(subrouter, app.RouteMiddleware[v])
		}
	}

	return registrar
}

func injectMiddleware(route *mux.Router, middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		route.Use(middleware)
	}
}

func getMethods(methods []string) []string {
	// check the length if 0
	// then default it to GET method
	if len(methods) == 0 {
		methods = []string{http.MethodGet}
	}

	return methods
}
