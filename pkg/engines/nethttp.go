package engines

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/urls"
)

type NetHttpEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request

	Response response.NetHttpResponse
	Request  request.NetHttpRequest
	Url      urls.NetHttpUrl
}

func NetHttp(w http.ResponseWriter, r *http.Request) *NetHttpEngine {
	res := response.NetHttp(w, r)
	url := urls.NetHttp(w, r)
	req := request.NetHttp(w, r, url)

	eg := NetHttpEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Response:       *res,
		Request:        *req,
		Url:            *url,
	}

	return &eg
}

func (m NetHttpEngine) GetRequest() interface{} {
	return m.Request
}

func (m NetHttpEngine) GetResponse() interface{} {
	return m.Response
}

func (m NetHttpEngine) GetUrl() interface{} {
	return m.Url
}
