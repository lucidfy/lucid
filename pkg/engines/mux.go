package engines

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/response"
)

type MuxEngine struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// ---

func (m MuxEngine) ParsedResponse() interface{} {
	return response.Mux(m.Writer, m.Request)
}

func (m MuxEngine) ParsedRequest() interface{} {
	return request.Mux(m.Writer, m.Request)
}

func (m MuxEngine) ParsedSession() interface{} {
	return session.Mux(m.Writer, m.Request)
}
