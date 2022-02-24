package session

import (
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

type Secured struct {
	SecuredCookie  *securecookie.SecureCookie
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Construct(w http.ResponseWriter, r *http.Request) *Secured {
	var sessionName string = os.Getenv("SESSION_NAME")

	gorvelSession, err := r.Cookie(sessionName)
	if err != nil && err == http.ErrNoCookie {
		return nil
	}

	return &Secured{
		SecuredCookie:  securecookie.New([]byte(gorvelSession.Value), nil),
		ResponseWriter: w,
		HttpRequest:    r,
	}
}

func (s *Secured) Set(name string, value string) (bool, error) {
	encoded, err := s.SecuredCookie.Encode(name, value)
	if err == nil {
		cookie := &http.Cookie{Name: name, Value: encoded, Path: "/"}
		http.SetCookie(s.ResponseWriter, cookie)
		return true, nil
	}
	return false, err
}

func (s *Secured) Get(name string) (*string, error) {
	if s.HttpRequest == nil {
		return nil, nil
	}

	cookie, err := s.HttpRequest.Cookie(name)
	if err == nil {
		var value string
		if err = s.SecuredCookie.Decode(name, cookie.Value, &value); err == nil {
			return &value, nil
		}
	}
	return nil, err
}
