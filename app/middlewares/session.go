package middlewares

import (
	"net/http"

	"github.com/daison12006013/lucid/pkg/facade/session"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session.GenerateSessionCookie(w, r)

		next.ServeHTTP(w, r)
	})
}
