package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
)

type ContextKey string

const ContextToken ContextKey = "csrf_token"

func CsrfProtectMiddleware(next http.Handler) http.Handler {
	protect := csrf.Protect(
		// 1st param is the csrf auth key
		[]byte(os.Getenv("CSRF_AUTH_KEY")),

		// 2nd param is the option with variadic param
		csrf.FieldName("csrf_token"),
	)
	return protect(next)
}

func CsrfSetterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := csrf.Token(r)
		w.Header().Set("X-CSRF-Token", token)

		next.ServeHTTP(w, r)
	})
}
