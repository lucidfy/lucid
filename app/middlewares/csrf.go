package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/helpers"
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
	opts := []csrf.Option{
		csrf.Path("/"),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.FieldName("csrf_token"),
		csrf.TrustedOrigins([]string{os.Getenv("CSRF_TRUSTED_ORIGIN")}),
		csrf.CookieName(helpers.Getenv("CSRF_NAME", "lucid_csrf")),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			trans := lang.Load(translations.Languages)
			engine := *engines.NetHttp(w, r, trans)
			handlers.HttpErrorHandler(engine, &errors.AppError{
				Message: "CSRF Failure",
				Error:   csrf.FailureReason(r),
				Code:    http.StatusForbidden,
			})
		})),
	}

	// for development purposes, we will allow the cookie to be
	// added in http only and should be secured only
	if helpers.Getenv("APP_ENV", "local") != "production" {
		opts = append(opts, csrf.HttpOnly(true), csrf.Secure(false))
	}

	key := []byte(helpers.Getenv("CSRF_KEY", os.Getenv("APP_KEY")))

	return csrf.Protect(key, opts...)(next)
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
