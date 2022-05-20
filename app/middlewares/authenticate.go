package middlewares

import (
	"net/http"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/session"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ses := session.File(w, r)

		_, app_err := ses.Get("authenticated")
		if app_err != nil {
			handlers.HttpErrorHandler(engines.Mux(w, r), &errors.AppError{
				Code:    http.StatusForbidden,
				Message: "Forbidden!",
				Error:   app_err.Error,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
