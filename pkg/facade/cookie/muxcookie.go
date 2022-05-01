package cookie

import (
	"net/http"
	"os"
	"strconv"

	"github.com/daison12006013/lucid/pkg/facade/crypt"
	"github.com/golang-module/carbon"
)

type MuxCookie struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxCookie {
	s := MuxCookie{
		ResponseWriter: w,
		HttpRequest:    r,
	}

	return &s
}

func (s *MuxCookie) CreateSessionCookie() interface{} {
	sessionKey := crypt.GenerateRandomString(20)
	s.Set(os.Getenv("SESSION_NAME"), sessionKey)
	return sessionKey
}

func (s *MuxCookie) Set(name string, value interface{}) (bool, error) {
	encoded, err := crypt.Encrypt(value)
	if err == nil {
		lifetime, err := strconv.Atoi(os.Getenv("SESSION_LIFETIME"))
		if err != nil {
			return false, err
		}
		cookie := &http.Cookie{
			Name:    name,
			Value:   encoded,
			Path:    "/",
			MaxAge:  lifetime,
			Expires: carbon.Now().AddSeconds(lifetime).Time,
			Domain:  os.Getenv("SESSION_DOMAIN"),
		}
		http.SetCookie(s.ResponseWriter, cookie)
		return true, nil
	}

	return false, err
}

func (s *MuxCookie) Get(name string) (interface{}, error) {
	if s.HttpRequest == nil {
		return nil, nil
	}

	cookie, err := s.HttpRequest.Cookie(name)
	if err != nil {
		return nil, err
	}

	decoded, err := crypt.Decrypt(cookie.Value)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (s *MuxCookie) Expire(name string) {
	cookie := &http.Cookie{Name: name, Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(s.ResponseWriter, cookie)
}
