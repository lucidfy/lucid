package engines

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/urls"
)

type MuxEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request

	Response response.MuxResponse
	Request  request.MuxRequest
	Url      urls.MuxUrl
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxEngine {
	res := response.Mux(w, r)
	url := urls.Mux(w, r)
	req := request.Mux(w, r, url)

	eg := MuxEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Response:       *res,
		Request:        *req,
		Url:            *url,
	}

	return &eg
}

func (m MuxEngine) GetRequest() interface{} {
	return m.Request
}

func (m MuxEngine) GetResponse() interface{} {
	return m.Response
}

func (m MuxEngine) GetUrl() interface{} {
	return m.Url
}
