package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/internal/facade/logger"
	"github.com/daison12006013/gorvel/internal/response"
)

func UserLists(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	records, err := users.Lists(1, 10, "id", "desc")
	if err != nil {
		logger.Error("handlers.UserLists error", err)
		panic(err)
	}

	response.View(w, "users/lists.html", map[string]interface{}{
		"records": records,
	})
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ! TODO
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ! TODO
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ! TODO
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ! TODO
}
