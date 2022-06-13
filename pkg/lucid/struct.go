package lucid

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterContract interface {
	Match(req *http.Request, match *mux.RouteMatch) bool
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Get(name string) *mux.Route
	GetRoute(name string) *mux.Route
	NewRoute() *mux.Route
	Name(name string) *mux.Route
	Handle(path string, handler http.Handler) *mux.Route
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	Headers(pairs ...string) *mux.Route
	Host(tpl string) *mux.Route
	MatcherFunc(f mux.MatcherFunc) *mux.Route
	Methods(methods ...string) *mux.Route
	Path(tpl string) *mux.Route
	PathPrefix(tpl string) *mux.Route
	Queries(pairs ...string) *mux.Route
	Schemes(schemes ...string) *mux.Route
	BuildVarsFunc(f mux.BuildVarsFunc) *mux.Route
	Walk(walkFn mux.WalkFunc) error
}
