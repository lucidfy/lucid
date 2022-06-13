package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucidfy/lucid/pkg/lucid"
)

func HttpAccessLogMiddleware(ctx lucid.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("Access on [%s]", r.Method), r.Host, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
