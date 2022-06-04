package urls

import (
	"net/http"
	"strings"

	"github.com/lucidfy/lucid/pkg/errors"
)

type NetHttpURL struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func NetHttp(w http.ResponseWriter, r *http.Request) *NetHttpURL {
	u := NetHttpURL{
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return &u
}

func (u *NetHttpURL) BaseURL() string {
	h := u.HttpRequest.URL.Host
	if len(h) > 0 {
		return h
	}

	// if URL.Host is empty, let's try to pull from request Host
	// then determine the scheme
	if len(u.HttpRequest.Host) > 0 {
		scheme := "http://"
		if strings.Contains(strings.ToLower(u.HttpRequest.Proto), "https") {
			scheme = "https://"
		}
		return scheme + u.HttpRequest.Host
	}

	//> if ever we can't resolve the Host from net/http, we can still base
	//> from our env config
	return BaseURL(nil)
}

func (u *NetHttpURL) CurrentURL() string {
	return u.BaseURL() + u.HttpRequest.URL.RequestURI()
}

func (u *NetHttpURL) PreviousURL() string {
	return u.HttpRequest.Referer()
}

func (u *NetHttpURL) RedirectPrevious() *errors.AppError {
	http.Redirect(u.ResponseWriter, u.HttpRequest, u.PreviousURL(), http.StatusFound)
	return nil
}
