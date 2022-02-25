package middlewares

import (
	"net/http"
	"os"
	"strconv"

	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/crypt"
	"github.com/golang-module/carbon"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsJsonRequest(w, r) {
			next.ServeHTTP(w, r)
			return
		}

		sessionName := os.Getenv("SESSION_NAME")
		_, err := r.Cookie(sessionName)

		if err != nil {
			switch err {
			case http.ErrNoCookie:
				randomStr := crypt.GenerateRandomString(24)
				lifetime, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
				if errors.Handler("error on getting SESSION_LIFETIME", err) {
					next.ServeHTTP(w, r)
					return
				}

				cookie := &http.Cookie{
					Name:    sessionName,
					Value:   randomStr,
					MaxAge:  lifetime,
					Expires: carbon.Now().AddSeconds(lifetime).Time,
					Path:    "/",
				}
				http.SetCookie(w, cookie)
			default:
				errors.Handler("error fetching Cookie "+sessionName, err)
			}
		}

		next.ServeHTTP(w, r)
	})
}
