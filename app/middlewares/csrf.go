package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
)

func CsrfProtectMiddleware(next http.Handler) http.Handler {
	protect := csrf.Protect(
		// 1st param is the csrf auth key
		[]byte(os.Getenv("CSRF_AUTH_KEY")),

		// 2nd param is the option with variadic param
		csrf.Path("/"),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.FieldName("csrf_token"),
		csrf.TrustedOrigins([]string{os.Getenv("CSRF_TRUSTED_ORIGIN")}),
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
