package session

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

type MuxSession struct {
	SecuredCookie  *securecookie.SecureCookie
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxSession {
	var sessionName string = os.Getenv("SESSION_NAME")

	gorvelSession, err := r.Cookie(sessionName)
	var securedCookie *securecookie.SecureCookie
	if err == nil {
		securedCookie = securecookie.New([]byte(gorvelSession.Value), nil)
	}

	s := MuxSession{
		SecuredCookie:  securedCookie,
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return &s
}

func (s *MuxSession) Set(name string, value interface{}) (bool, error) {
	if s.SecuredCookie == nil {
		return false, fmt.Errorf("SecuredCookie is empty")
	}

	encoded, err := s.SecuredCookie.Encode(name, &value)
	if err == nil {
		cookie := &http.Cookie{Name: name, Value: encoded, Path: "/"}
		http.SetCookie(s.ResponseWriter, cookie)
		return true, nil
	}
	return false, err
}

func (s *MuxSession) Get(name string) (interface{}, error) {
	if s.HttpRequest == nil {
		return nil, nil
	}

	if s.SecuredCookie == nil {
		return nil, fmt.Errorf("SecuredCookie is empty")
	}

	cookie, err := s.HttpRequest.Cookie(name)
	if err == nil {
		var value interface{}
		if err = s.SecuredCookie.Decode(name, cookie.Value, &value); err == nil {
			return &value, nil
		}
	}
	return nil, err
}

func (s *MuxSession) SetFlash(name string, value interface{}) {
	name = "flash-" + name
	s.Set(name, value)
}

func (s *MuxSession) GetFlash(name string) interface{} {
	name = "flash-" + name
	value, err := s.Get(name)
	if (err != nil && err == http.ErrNoCookie) || value == nil {
		return nil
	}
	// delete the cookie by expiring it immediately!
	deleteCookie := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0), Path: "/"}
	http.SetCookie(s.ResponseWriter, deleteCookie)
	return value
}
