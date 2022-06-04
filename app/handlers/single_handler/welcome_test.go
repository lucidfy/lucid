package single_handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/lucid"
	"github.com/lucidfy/lucid/resources/translations"
)

func init() {
	env.LoadEnv()
}

func TestWelcome(t *testing.T) {
	// initializes the engine router
	path := WelcomeRoute.Path
	handler := WelcomeRoute.Handler
	method := WelcomeRoute.Method[0]
	handler_func := func(w http.ResponseWriter, r *http.Request) {
		engine := *engines.NetHttp(w, r, lang.Load(translations.Languages))
		ctx := context.Background()
		ctx = context.WithValue(ctx, lucid.EngineCtx{}, engine)
		app_err := handler(lucid.NewContext(ctx))
		ctx.Done()

		if app_err != nil {
			handlers.HttpErrorHandler(engine, app_err, nil)
		}
	}

	// httptest
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	handler_func(w, req)
	result := w.Result()

	// test the status code

	if status := result.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// test the content

	b, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	content := string(b)

	has_welcome_lucid := strings.Contains(content, `Welcome to <a href="https://github.com/lucidfy/lucid" class="font-light">Lucid</a>`)
	if !has_welcome_lucid {
		t.Error("content should have the Welcome to Lucid")
	}
}
