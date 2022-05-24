package routes

import (
	"context"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type FiberRoutes struct {
	App               *fiber.App
	GlobalMiddlewares []interface{}
	RouteMiddlewares  map[string]interface{}
	HttpErrorHandler  func(engines.EngineContract, *errors.AppError)
	Translation       *lang.Translations
}

func Fiber(t *lang.Translations) *FiberRoutes {
	return &FiberRoutes{Translation: t}
}

func (fr *FiberRoutes) AddGlobalMiddlewares(base []interface{}) *FiberRoutes {
	fr.GlobalMiddlewares = base
	return fr
}

func (fr *FiberRoutes) AddRouteMiddlewares(base map[string]interface{}) *FiberRoutes {
	fr.RouteMiddlewares = base
	return fr
}

// Here, you can find how we iterate the routes() function,
// we're using gorilla/mux package to serve our routing with
// extensive support with http requests + middlewares.
func (fr *FiberRoutes) Register(routings *[]Routing) *FiberRoutes {
	fr.App = fiber.New()

	for _, routing := range *fr.Explain(routings) {
		fr.register(routing)
	}

	return fr
}

func (fr FiberRoutes) Explain(base *[]Routing) *[]Routing {
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

func (fr *FiberRoutes) register(route Routing) {
	// serve static
	if len(route.Static) != 0 {
		fr.App.Static(route.Path, route.Static)
		return
	}

	// collate all middlewares into slice
	var mids []func(*fiber.Ctx) error
	if route.WithGlobalMiddleware == nil || route.WithGlobalMiddleware == true {
		for _, v := range fr.GlobalMiddlewares {
			mids = append(mids, adaptor.HTTPMiddleware(
				v.(func(http.Handler) http.Handler),
			))
		}
	}
	for _, v := range route.Middlewares {
		mids = append(mids, adaptor.HTTPMiddleware(
			fr.RouteMiddlewares[v].(func(http.Handler) http.Handler),
		))
	}

	// get handler
	fiber_handle := adaptorToHttpHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.NetHttp(w, r, fr.Translation)
		e := route.Handler(engine)
		if e != nil {
			fr.HttpErrorHandler(engine, e)
		}
	})

	// put all middlewares and get the router
	router := fr.App.Add("USE", route.Path, mids...)

	// now the final sauce, we register those methods now
	for _, method := range getMethods(route.Method) {
		router.Add(method, route.Path, fiber_handle)
	}
}

type netHTTPResponseWriter struct {
	statusCode int
	h          http.Header
	body       []byte
}

func (w *netHTTPResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *netHTTPResponseWriter) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *netHTTPResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *netHTTPResponseWriter) Write(p []byte) (int, error) {
	w.body = append(w.body, p...)
	return len(p), nil
}

func adaptorToHttpHandlerFunc(h http.HandlerFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := newFasthttpHandler(h, c.AllParams())
		handler(c.Context())
		return nil
	}
}

func newFasthttpHandler(h http.Handler, vars map[string]string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var r http.Request
		if err := fasthttpadaptor.ConvertRequest(ctx, &r, true); err != nil {
			ctx.Logger().Printf("cannot parse requestURI %q: %v", r.RequestURI, err)
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			return
		}

		var w netHTTPResponseWriter
		r = *r.WithContext(context.WithValue(ctx, request.VarsKey, vars))
		h.ServeHTTP(&w, &r)

		ctx.SetStatusCode(w.StatusCode())
		haveContentType := false
		for k, vv := range w.Header() {
			if k == fasthttp.HeaderContentType {
				haveContentType = true
			}

			for _, v := range vv {
				ctx.Response.Header.Add(k, v)
			}
		}
		if !haveContentType {
			// From net/http.ResponseWriter.Write:
			// If the Header does not contain a Content-Type line, Write adds a Content-Type set
			// to the result of passing the initial 512 bytes of written data to DetectContentType.
			l := 512
			if len(w.body) < 512 {
				l = len(w.body)
			}
			ctx.Response.Header.Set(fasthttp.HeaderContentType, http.DetectContentType(w.body[:l]))
		}
		ctx.Write(w.body) //nolint:errcheck
	}
}
