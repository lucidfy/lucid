package session

import (
	"net/http"
	"os"
	"strconv"

	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/crypt"
	"github.com/golang-module/carbon"
)

type SesionContract interface {
	Set(name string, value string) (bool, error)
	Get(name string) (*string, error)
	SetFlash(name string, value string)
	GetFlash(name string) *string
	SetFlashMap(name string, values interface{})
	GetFlashMap(name string) *map[string]interface{}
}

func GenerateSessionCookie(w http.ResponseWriter, r *http.Request) error {
	sessionName := os.Getenv("SESSION_NAME")
	_, err := r.Cookie(sessionName)

	if err != nil {
		switch err {
		case http.ErrNoCookie:
			randomString := crypt.GenerateRandomString(128)
			lifetime, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
			if errors.Handler("error on getting SESSION_LIFETIME", err) {
				return err
			}

			cookie := &http.Cookie{
				Name:    sessionName,
				Value:   randomString,
				MaxAge:  lifetime,
				Expires: carbon.Now().AddSeconds(lifetime).Time,
				Path:    "/",
			}
			http.SetCookie(w, cookie)
		default:
			errors.Handler("error fetching Cookie "+sessionName, err)
			return err
		}
	}

	return nil
}
