package middlewares

import (
	"errors"
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/cookie"
	"github.com/lucidfy/lucid/pkg/helpers"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coo := cookie.New(w, r)
		_, app_err := coo.Get(helpers.Getenv("SESSION_NAME", "lucid_session"))

		if app_err != nil && errors.Is(app_err.Error, http.ErrNoCookie) {
			coo.CreateSessionCookie()
		}

		next.ServeHTTP(w, r)
	})
}
