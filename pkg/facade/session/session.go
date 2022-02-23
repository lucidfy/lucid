package session

import (
	"net/http"
	"os"

	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/gorilla/securecookie"
)

type Secured struct {
	SecuredCookie  *securecookie.SecureCookie
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Construct(w http.ResponseWriter, r *http.Request) Secured {
	var sessionName string = os.Getenv("SESSION_NAME")

	gorvelSession, err := r.Cookie(sessionName)
	if errors.Handler("error fetching Cookie "+sessionName, err) {
		panic(err)
	}

	return Secured{
		SecuredCookie:  securecookie.New([]byte(gorvelSession.Value), nil),
		ResponseWriter: w,
		HttpRequest:    r,
	}
}

func (s Secured) Set(name string, value string) (bool, error) {
	encoded, err := s.SecuredCookie.Encode(name, value)
	if err == nil {
		cookie := &http.Cookie{Name: name, Value: encoded, Path: "/"}
		http.SetCookie(s.ResponseWriter, cookie)
		return true, nil
	}
	return false, err
}

func (s Secured) Get(name string) (*string, error) {
	cookie, err := s.HttpRequest.Cookie(name)
	if err == nil {
		var value string
		if err = s.SecuredCookie.Decode(name, cookie.Value, &value); err == nil {
			return &value, nil
		}
	}
	return nil, err
}
