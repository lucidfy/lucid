package routes

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

// interface

type RouteInterface interface {
	Register(base *[]Routing) interface{}
	Explain(base *[]Routing) interface{}
}

// structs

// type AppError struct {
// 	Error   error
// 	Message interface{}
// 	Code    interface{}
// }
type Middlewares []string
type Queries []string
type Method []string
type Resources map[string]Handler
type Handler func(engines.EngineInterface) *errors.AppError
type Routing struct {
	Name        string
	Path        string
	Method      []string
	Queries     Queries
	Handler     Handler
	Resources   map[string]Handler
	Middlewares []string
}

// helpers

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
			routings = append(
				routings,
				Routing{
					Path:        route.Path + "/{id}",
					Handler:     handler,
					Method:      Method{"PUT"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update",
				},
				Routing{
					Path:        route.Path + "/{id}/update",
					Handler:     handler,
					Method:      Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update.via-post",
				},
			)
		case "destroy":
			routings = append(
				routings,
				Routing{
					Path:        route.Path + "/{id}",
					Handler:     handler,
					Method:      Method{"DELETE"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy",
				},
				Routing{
					Path:        route.Path + "/{id}/delete",
					Handler:     handler,
					Method:      Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy.via-post",
				},
			)
		}
	}
	return routings
}

func getMethods(methods Method) Method {
	if len(methods) == 0 {
		methods = Method{http.MethodGet}
	}
	return methods
}
