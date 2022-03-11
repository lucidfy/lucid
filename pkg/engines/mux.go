package engines

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/facade/response"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/facade/urls"
)

type MuxEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request

	Response response.MuxResponse
	Request  request.MuxRequest
	Session  session.MuxSession
	Url      urls.MuxUrl
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxEngine {
	res := *response.Mux(w, r)
	url := urls.Mux(w, r)
	ses := session.Mux(w, r)
	req := *request.Mux(w, r, url, ses)

	eg := MuxEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Response:       res,
		Request:        req,
		Session:        *ses,
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

func (m MuxEngine) GetSession() interface{} {
	return m.Session
}

func (m MuxEngine) GetUrl() interface{} {
	return m.Url
}
