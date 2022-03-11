package urls

import "net/http"

type MuxUrl struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxUrl {
	u := MuxUrl{
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return &u
}

func (u *MuxUrl) CurrentUrl() string {
	h := u.HttpRequest.URL.Host
	if len(h) > 0 {
		return h
	}

	//> if ever we can't resolve the Host from net/http, we can still base
	//> from our env config
	return BaseUrl(nil)
}

func (u *MuxUrl) FullUrl() string {
	return u.CurrentUrl() + u.HttpRequest.URL.RequestURI()
}

func (u *MuxUrl) PreviousUrl() string {
	return u.HttpRequest.Referer()
}

func (u *MuxUrl) RedirectPrevious() {
	http.Redirect(u.ResponseWriter, u.HttpRequest, u.PreviousUrl(), http.StatusFound)
}
