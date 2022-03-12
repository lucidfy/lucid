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
	eg := MuxEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Response:       *response.Mux(w, r),
		Request:        *request.Mux(w, r),
		Session:        *session.Mux(w, r),
		Url:            *urls.Mux(w, r),
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
