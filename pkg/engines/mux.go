package engines

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/response"
)

type MuxEngine struct {
	HttpResponseWriter http.ResponseWriter
	HttpRequest        *http.Request

	Response response.MuxResponse
	Request  request.MuxRequest
}

func Mux(w http.ResponseWriter, r *http.Request) MuxEngine {
	return MuxEngine{
		HttpResponseWriter: w,
		HttpRequest:        r,
		Response:           response.Mux(w, r),
		Request:            request.Mux(w, r),
	}
}

func (m MuxEngine) ParsedResponse() interface{} {
	return m.Response
}

func (m MuxEngine) ParsedRequest() interface{} {
	return m.Request
}

func (m MuxEngine) ParsedSession() interface{} {
	return session.Mux(m.HttpResponseWriter, m.HttpRequest)
}
