package loader

import (
	"net/http"
	"strings"

	"github.com/lucidfy/lucid/pkg/facade/routes"
)

type LoaderInterface interface {
	Register(base *[]routes.Routing) interface{}
	Explain(base *[]routes.Routing) interface{}
}

func stripDoubleSlash(str string) string {
	return strings.Replace(str, "//", "/", -1)
}

func resources(route routes.Routing) []routes.Routing {
	routings := []routes.Routing{}
	id_regex := "{id:[0-9]+}"

	for action, handler := range route.Resources {
		switch action {
		case "index":
			routings = append(routings, routes.Routing{
				Path:        stripDoubleSlash(route.Path),
				Handler:     handler,
				Method:      routes.Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".lists",
			})
		case "create":
			routings = append(routings, routes.Routing{
				Path:        stripDoubleSlash(route.Path + "/create"),
				Handler:     handler,
				Method:      routes.Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".create",
			})
		case "store":
			routings = append(routings, routes.Routing{
				Path:        stripDoubleSlash(route.Path),
				Handler:     handler,
				Method:      routes.Method{"POST"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".store",
			})
		case "show":
			routings = append(routings, routes.Routing{
				Path:        stripDoubleSlash(route.Path + "/" + id_regex),
				Handler:     handler,
				Method:      routes.Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".show",
			})
		case "edit":
			routings = append(routings, routes.Routing{
				Path:        stripDoubleSlash(route.Path + "/" + id_regex + "/edit"),
				Handler:     handler,
				Method:      routes.Method{"GET"},
				Middlewares: route.Middlewares,
				Name:        route.Name + ".edit",
			})
		case "update":
			routings = append(
				routings,
				routes.Routing{
					Path:        stripDoubleSlash(route.Path + "/" + id_regex),
					Handler:     handler,
					Method:      routes.Method{"PUT"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update",
				},
				routes.Routing{
					Path:        stripDoubleSlash(route.Path + "/" + id_regex + "/update"),
					Handler:     handler,
					Method:      routes.Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".update.via-post",
				},
			)
		case "destroy":
			routings = append(
				routings,
				routes.Routing{
					Path:        stripDoubleSlash(route.Path + "/" + id_regex),
					Handler:     handler,
					Method:      routes.Method{"DELETE"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy",
				},
				routes.Routing{
					Path:        stripDoubleSlash(route.Path + "/" + id_regex + "/delete"),
					Handler:     handler,
					Method:      routes.Method{"POST"},
					Middlewares: route.Middlewares,
					Name:        route.Name + ".destroy.via-post",
				},
			)
		}
	}
	return routings
}

func getMethods(methods routes.Method) routes.Method {
	if len(methods) == 0 {
		methods = routes.Method{http.MethodGet}
	}
	return methods
}
