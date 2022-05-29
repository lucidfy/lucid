package middlewares

import (
	"net/http"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/resources/translations"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := lang.Load(translations.Languages)
		engine := engines.NetHttp(w, r, t)
		ses := engine.Session

		if ses != nil {
			_, app_err := ses.Get("authenticated")
			if app_err != nil {
				t := lang.Load(translations.Languages)
				handlers.HttpErrorHandler(engines.NetHttp(w, r, t), &errors.AppError{
					Code:    http.StatusForbidden,
					Message: "Forbidden!",
					Error:   app_err.Error,
				}, nil)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
