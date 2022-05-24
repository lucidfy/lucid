package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/resources/translations"
)

// CsrfShouldSkipMiddleware here, we determine if we should skip the csrf
// mainly we skip if the condition inside IsJsonRequest(...)
// returns true, basically if the request wanted a json response
func CsrfShouldSkipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is_json := func(w http.ResponseWriter, r *http.Request) (is_json bool) {
			t := lang.Load(translations.Languages)
			rp := request.NetHttp(w, r, t, nil)
			is_json = rp.IsJson() || rp.WantsJson()
			return
		}(w, r)

		if is_json {
			r = csrf.UnsafeSkipCheck(r)
		}

		next.ServeHTTP(w, r)
	})
}

// CsrfProtectMiddleware here, we initialize gorilla's csrf
// by default we set the csrf_auth_key too
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

// CsrfSetterMiddleware here, we pass the token as X-CSRF-Token via header
func CsrfSetterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := csrf.Token(r)
		if len(token) > 0 {
			w.Header().Set("X-CSRF-Token", token)
		}
		next.ServeHTTP(w, r)
	})
}
