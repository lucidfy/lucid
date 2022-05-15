package middlewares

import (
	"errors"
	"net/http"
	"os"

	"github.com/lucidfy/lucid/pkg/facade/cookie"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coo := cookie.New(w, r)
		_, err := coo.Get(os.Getenv("SESSION_NAME"))

		if err != nil && errors.Is(err, http.ErrNoCookie) {
			coo.CreateSessionCookie()
		}

		next.ServeHTTP(w, r)
	})
}
