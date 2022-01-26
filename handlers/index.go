package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/glade"
	"github.com/daison12006013/gorvel/structs"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := structs.HomepageData{
		Title:  "Gorvel Rocks!",
		Params: mux.Vars(r),
	}

	glade.View(w, "views/index.html", data)
}
