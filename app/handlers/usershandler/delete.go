package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	rp := request.Parse(w, r)

	// prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	id := rp.GetFirst("id", nil)

	exists, _ := users.Exists(id)
	if exists {
		users.Delete(id)
	} else {
		message = "Record not found!"
		status = http.StatusNotFound
	}

	// prepare the data
	if rp.IsJson() && rp.WantsJson() {
		response.Json(w, map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
	}

	rp.SetFlash("success", message).RedirectPrevious()
}
