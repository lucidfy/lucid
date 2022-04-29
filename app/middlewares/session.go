package middlewares

import (
	"net/http"
	"os"
	"strconv"

	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/crypt"
	"github.com/golang-module/carbon"
)

func SessionPersistenceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionName := os.Getenv("SESSION_NAME")
		_, err := r.Cookie(sessionName)

		if err != nil {
			switch err {
			case http.ErrNoCookie:
				sessionRandomPassword := crypt.GenerateRandomString(128)
				lifetime, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
				if errors.Handler("error on getting SESSION_LIFETIME", err) {
					next.ServeHTTP(w, r)
					return
				}

				cookie := &http.Cookie{
					Name:    sessionName,
					Value:   sessionRandomPassword,
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
