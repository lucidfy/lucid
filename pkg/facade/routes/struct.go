package routes

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/lucid"
	"github.com/lucidfy/lucid/resources/translations"
)

type Middlewares []string
type Queries []string
type Method []string
type Resources map[string]Handler
type Handler func(lucid.Context) *errors.AppError
type Routing struct {
	Name                 string
	Path                 string
	Prefix               bool
	Method               []string
	Queries              Queries
	Handler              Handler
	Resources            map[string]Handler
	Middlewares          []string
	Static               string
	WithGlobalMiddleware interface{}
}

type RoutingTest struct {
	ResponseRecorder *httptest.ResponseRecorder
	Request          *http.Request
	Err              error
	Testing          *testing.T
	Content          string
	Routing          Routing
}

func (r Routing) LoadTester(res *httptest.ResponseRecorder, req *http.Request) *RoutingTest {
	path := r.Path
	method := r.Method[0]

	// check response record if nil then initialize it
	if res == nil {
		res = httptest.NewRecorder()
	}

	// make a new request basing from the handler's method and path
	// and get the body provided from param
	var err error
	if req == nil {
		req, err = http.NewRequest(method, path, nil)
	}

	return &RoutingTest{
		ResponseRecorder: res,
		Request:          req,
		Err:              err,
		Routing:          r,
	}
}

func (rt *RoutingTest) AssertUsing(t *testing.T) {
	rt.Testing = t
}

func (rt RoutingTest) AssertStatus(expectation uint) {
	result := rt.ResponseRecorder.Result()

	if status := result.StatusCode; status != int(expectation) {
		rt.Testing.Errorf("status: got [[%d]] expects [[%d]]", result.StatusCode, expectation)
	}
}

func (rt *RoutingTest) CallHandler() {
	handler := rt.Routing.Handler
	handler_func := func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.NetHttp(w, r, lang.Load(translations.Languages))
		ctx := context.Background()
		ctx = context.WithValue(ctx, lucid.EngineCtx{}, engine)
		app_err := handler(lucid.New(ctx))
		ctx.Done()

		if app_err != nil {
			handlers.HttpErrorHandler(engine, app_err, nil)
		}
	}

	// now bootstrap the handler
	handler_func(rt.ResponseRecorder, rt.Request)

	result := rt.ResponseRecorder.Result()
	b, err := io.ReadAll(result.Body)
	if err != nil {
		rt.Testing.Fatal(err)
	}
	rt.Content = string(b)
}

func (rt RoutingTest) AssertResponseContains(expectation string) {
	content := rt.Content
	if has := strings.Contains(content, expectation); !has {
		rt.Testing.Errorf("content: got [[%s]] does not contain with [[%s]]", content, expectation)
	}
}
