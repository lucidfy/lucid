package middlewares

import (
	"net/http"
	"os"

	"github.com/daison12006013/lucid/pkg/facade/cookie"
	"github.com/daison12006013/lucid/pkg/facade/crypt"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("SESSION_NAME")
		coo := cookie.Mux(w, r)
		_, err := coo.Get(name)

		if err != nil {
			switch err {
			case http.ErrNoCookie:
				coo.Set(name, crypt.GenerateRandomString(20))
			}
		}

		next.ServeHTTP(w, r)
	})
}
