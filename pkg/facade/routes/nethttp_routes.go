package routes

import "github.com/gorilla/mux"

type NetHttpRouter struct {
	Route *mux.Router
}

func NetHttp(r *mux.Router) *NetHttpRouter {
	return &NetHttpRouter{
		Route: r,
	}
}
