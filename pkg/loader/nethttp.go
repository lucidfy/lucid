package loader

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/lucid"
)

type NetHttpLoader struct {
	Router            *mux.Router
	Routings          *[]routes.Routing
	GlobalMiddlewares []interface{}
	RouteMiddlewares  map[string]interface{}
	HttpErrorHandler  func(engines.EngineContract, *errors.AppError, interface{})
	Translation       *lang.Translations
}

func NetHttp(t *lang.Translations) *NetHttpLoader {
	mr := &NetHttpLoader{
		Translation: t,
	}
	mr.Router = mux.NewRouter().StrictSlash(true)
	return mr
}

func (mr *NetHttpLoader) AddGlobalMiddlewares(base []interface{}) *NetHttpLoader {
	mr.GlobalMiddlewares = base
	return mr
}

func (mr *NetHttpLoader) AddRouteMiddlewares(base map[string]interface{}) *NetHttpLoader {
	mr.RouteMiddlewares = base
	return mr
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func (mr NetHttpLoader) Register(base *[]routes.Routing) *mux.Router {
	// each routing should be interpreted as subrouter
	// the subrouter in mux isolates each path with
	// a way to register a repetitive middlewares
	for _, routing := range *mr.Explain(base) {
		submr := mr
		submr.Router = mr.Router.NewRoute().Subrouter()
		submr.register(routing)
	}

	return mr.Router
}

func (mr *NetHttpLoader) Explain(base *[]routes.Routing) *[]routes.Routing {
	routings := []routes.Routing{}
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

func (mr *NetHttpLoader) register(route routes.Routing) {
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
		engine := *engines.NetHttp(w, r, mr.Translation)
		engine.HttpErrorHandler = mr.HttpErrorHandler

		ctx := context.Background()
		ctx = context.WithValue(ctx, lucid.EngineCtx{}, engine)
		ctx = context.WithValue(ctx, lucid.RouterCtx{}, mr.Router)
		e := route.Handler(ctx)
		ctx.Done()

		if e != nil {
			mr.HttpErrorHandler(engine, e, nil)
		}

		r.Body.Close()
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

func (mr *NetHttpLoader) routeUse(middlewares ...mux.MiddlewareFunc) {
	for _, middleware := range middlewares {
		mr.Router.Use(middleware)
	}
}
