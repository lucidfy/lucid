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
}

func (r Routing) TestLoad(body io.Reader) RoutingTest {
	path := r.Path
	handler := r.Handler
	method := r.Method[0]

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

	// httptest
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, body)
	handler_func(rr, req)

	return RoutingTest{
		ResponseRecorder: rr,
		Request:          req,
		Err:              err,
	}
}

func (r *RoutingTest) Uses(t *testing.T) {
	r.Testing = t
}

func (r RoutingTest) AssertStatus(expectation uint) {
	result := r.ResponseRecorder.Result()

	if status := result.StatusCode; status != int(expectation) {
		r.Testing.Errorf("status: got [[%d]] expects [[%d]]", result.StatusCode, expectation)
	}
}

func (r RoutingTest) AssertResponseContains(expectation string) {
	result := r.ResponseRecorder.Result()

	b, err := io.ReadAll(result.Body)
	if err != nil {
		r.Testing.Fatal(err)
	}
	content := string(b)

	if has := strings.Contains(content, expectation); !has {
		r.Testing.Errorf("content: got [[%s]] does not contain with [[%s]]", content, expectation)
	}
}
