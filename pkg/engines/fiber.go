package engines

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/response"
)

// TODO: replace this with proper fiber implementation!

type FiberEngine struct {
	HttpResponseWriter http.ResponseWriter
	HttpRequest        *http.Request

	Response response.MuxResponse
	Request  request.MuxRequest
}

// ---

func Fiber(w http.ResponseWriter, r *http.Request) FiberEngine {
	return FiberEngine{
		HttpResponseWriter: w,
		HttpRequest:        r,
		Response:           response.Mux(w, r),
		Request:            request.Mux(w, r),
	}
}

func (m FiberEngine) ParsedResponse() interface{} {
	return m.Response
}

func (m FiberEngine) ParsedRequest() interface{} {
	return m.Request
}

func (m FiberEngine) ParsedSession() interface{} {
	return session.Mux(m.HttpResponseWriter, m.HttpRequest)
}
