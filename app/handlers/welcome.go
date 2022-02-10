package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/internal/response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	response.View(w, "index.html", map[string]interface{}{
		"title": "Gorvel Rocks!",
	})
}
