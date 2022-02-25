package engines

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/response"
)

// TODO: replace this with proper fiber implementation!

type FiberEngine struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// ---

func (f FiberEngine) ParsedResponse() interface{} {
	return response.Mux(f.Writer, f.Request)
}

func (f FiberEngine) ParsedRequest() interface{} {
	return request.Mux(f.Writer, f.Request)
}

func (f FiberEngine) ParsedSession() interface{} {
	return session.Mux(f.Writer, f.Request)
}
