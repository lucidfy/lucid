package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
)

type NetHttpRoutes struct {
	Router            *mux.Router
	Routings          *[]Routing
	GlobalMiddlewares []interface{}
	RouteMiddlewares  map[string]interface{}
	HttpErrorHandler  func(engines.EngineContract, *errors.AppError)
	Translation       *lang.Translations
}

func NetHttp(t *lang.Translations) *NetHttpRoutes {
	mr := &NetHttpRoutes{
		Translation: t,
	}
	mr.Router = mux.NewRouter().StrictSlash(true)
	return mr
}

func (mr *NetHttpRoutes) AddGlobalMiddlewares(base []interface{}) *NetHttpRoutes {
	mr.GlobalMiddlewares = base
	return mr
}

func (mr *NetHttpRoutes) AddRouteMiddlewares(base map[string]interface{}) *NetHttpRoutes {
	mr.RouteMiddlewares = base
	return mr
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func (mr *NetHttpRoutes) Register(base *[]Routing) *mux.Router {
	// each routing should be interpreted as subrouter
	// the subrouter in mux isolates each path with
	// a way to register a repetitive middlewares
	for _, routing := range *mr.Explain(base) {
		submr := NetHttp(mr.Translation).
			AddGlobalMiddlewares(mr.GlobalMiddlewares).
			AddRouteMiddlewares(mr.RouteMiddlewares)

		submr.Router = mr.Router.NewRoute().Subrouter()
		submr.register(routing)
	}

	return mr.Router
}

func (mr *NetHttpRoutes) Explain(base *[]Routing) *[]Routing {
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

func (mr *NetHttpRoutes) register(route Routing) {
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
		defer r.Body.Close()

		engine := *engines.NetHttp(w, r, mr.Translation)
		e := route.Handler(engine)
		if e != nil {
			mr.HttpErrorHandler(engine, e)
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

	if route.WithGlobalMiddleware == nil || route.WithGlobalMiddleware == true {
		for _, mid := range mr.GlobalMiddlewares {
			mr.routeUse(mid.(func(http.Handler) http.Handler))
		}
	}

	mids := make(map[string]func(http.Handler) http.Handler)
	for k, mid := range mr.RouteMiddlewares {
		mids[k] = mid.(func(http.Handler) http.Handler)
	}

	for _, v := range route.Middlewares {
		mr.routeUse(mids[v])
	}
}

func (mr *NetHttpRoutes) routeUse(middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		mr.Router.Use(middleware)
	}
}
