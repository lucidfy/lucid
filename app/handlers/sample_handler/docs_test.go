package sample_handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/lucid"
	"github.com/lucidfy/lucid/resources/translations"
)

func init() {
	env.LoadEnvForTests()
}

func responseRecord(handler routes.Handler, method string, path string, body io.Reader) (*httptest.ResponseRecorder, *http.Request, error) {
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
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, body)
	handler_func(w, req)

	return w, req, err
}

func TestDocs(t *testing.T) {
	rt := DocsRoute.LoadTester(nil, nil)
	rt.CallHandler()

	rt.AssertUsing(t)
	rt.AssertStatus(200)
	rt.AssertResponseContains("# Installation")
}
