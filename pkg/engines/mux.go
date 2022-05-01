package engines

import (
	"net/http"

	"github.com/daison12006013/lucid/pkg/facade/cookie"
	"github.com/daison12006013/lucid/pkg/facade/request"
	"github.com/daison12006013/lucid/pkg/facade/response"
	"github.com/daison12006013/lucid/pkg/facade/urls"
)

type MuxEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request

	Response response.MuxResponse
	Request  request.MuxRequest
	Cookie   cookie.MuxCookie
	Url      urls.MuxUrl
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxEngine {
	res := response.Mux(w, r)
	url := urls.Mux(w, r)
	coo := cookie.Mux(w, r)
	req := request.Mux(w, r, url)

	eg := MuxEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Response:       *res,
		Request:        *req,
		Cookie:         *coo,
		Url:            *url,
	}

	return &eg
}

func (m MuxEngine) GetResponse() interface{} {
	return m.Response
}

func (m MuxEngine) GetRequest() interface{} {
	return m.Request
}

func (m MuxEngine) GetCookie() interface{} {
	return m.Cookie
}

func (m MuxEngine) GetUrl() interface{} {
	return m.Url
}
