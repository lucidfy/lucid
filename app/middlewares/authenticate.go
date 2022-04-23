package middlewares

import (
	"net/http"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// engine := engines.Mux(w, r)
		// ses := engine.Session

		// x, err := ses.Get("user")
		// if err != nil {
		// 	handlers.HttpErrorHandler(*engine, &errors.AppError{
		// 		Code:    http.StatusForbidden,
		// 		Message: "Forbidden!",
		// 		Error:   err,
		// 	})
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
