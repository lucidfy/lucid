package config

import (
	"github.com/daison12006013/gorvel/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	h := mux.NewRouter()

	h.HandleFunc("/", handlers.Home)

	return h
}
