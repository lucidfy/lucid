package routes

import (
	"net/http"

	"github.com/daison12006013/gorvel/app"
	"github.com/daison12006013/gorvel/app/handlers"
	"github.com/daison12006013/gorvel/app/handlers/usershandler"
	"github.com/gorilla/mux"
)

func Routes() *[]routeStruct {
	l := &[]routeStruct{
		{
			path:    "/",
			method:  []string{"GET"},
			handler: handlers.Home,
		},
		{
			path: "/users",
			resources: map[string]handlerStruct{
				"create": usershandler.Create, //   POST   /users
				"delete": usershandler.Delete, //   DELETE /users/{id?}
				"find":   usershandler.Find,   //   GET    /users/{id}
				"lists":  usershandler.Lists,  //   GET    /users
				"update": usershandler.Update, //   PUT    /users
			},
			middlewares: []string{"auth"},
		},
	}

	return l
}

// ---------------------------------------------------------------------

type handlerStruct func(http.ResponseWriter, *http.Request)
type routeStruct struct {
	path        string
	method      []string
	queries     []string
	handler     handlerStruct
	resources   map[string]handlerStruct
	middlewares []string
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func Register() *mux.Router {
	registrar := mux.NewRouter().StrictSlash(true)

	// Register the global middlewares
	appendMiddlewares(registrar, app.Middleware...)

	for _, route := range *Routes() {
		subrouter := registrar.NewRoute().Subrouter()

		if len(route.resources) > 0 {
			resources(subrouter, route)
		} else {
			register(subrouter, route.path, route.handler, route.method, route)
		}
	}

	return registrar
}

func resources(router *mux.Router, route routeStruct) {
	for action, handler := range route.resources {
		switch action {
		case "create":
			register(router, route.path, handler, []string{"POST"}, route)
		case "delete":
			register(router, route.path, handler, []string{"DELETE"}, route)
		case "find":
			register(router, route.path+"/{id}", handler, []string{"GET"}, route)
		case "lists":
			register(router, route.path, handler, []string{"GET"}, route)
		case "update":
			register(router, route.path+"/{id}", handler, []string{"PUT"}, route)
		}
	}
}

func register(
	subrouter *mux.Router,
	routePath string,
	routeHandler handlerStruct,
	routeMethod []string,
	route routeStruct,
) {
	subrouter.HandleFunc(routePath, routeHandler).
		Methods(getMethods(routeMethod)...).
		Queries(route.queries...)

	for _, v := range route.middlewares {
		appendMiddlewares(subrouter, app.RouteMiddleware[v])
	}
}

func appendMiddlewares(route *mux.Router, middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		route.Use(middleware)
	}
}

func getMethods(methods []string) []string {
	if len(methods) == 0 {
		methods = []string{http.MethodGet}
	}
	return methods
}
