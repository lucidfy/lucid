package middlewares

import (
	"net/http"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate the request was authenticated
		next.ServeHTTP(w, r)
	})
}
