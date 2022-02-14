package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	req := request.Parse(r)

	// prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	// check if the record exists
	exists := users.Exists(*req.Input("id"))

	// if exists, delete the record
	if exists {
		ok := users.DeleteById(*req.Input("id"))
		if !ok {
			message = "Cannot delete record!"
			status = http.StatusInternalServerError
		}
	} else {
		message = "Record not found!"
		status = http.StatusNotFound
	}

	// prepare the data
	response.Json(w, map[string]interface{}{
		"message": message,
	}, status)
}
