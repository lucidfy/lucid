package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/response"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Create Form",
	}

	response.View(
		w,
		[]string{"base.go.html", "users/create.go.html"},
		data,
	)
}

func Store(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	// req := request.Parse(w, r)

	// prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	// * TODO

	// prepare the data
	response.Json(w, map[string]interface{}{
		"message": message,
	}, status)
}
