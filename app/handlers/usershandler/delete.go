package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	engine := engines.MuxEngine{Writer: w, Request: r}
	request := engine.ParsedRequest().(request.MuxRequest)
	response := engine.ParsedResponse().(response.MuxResponse)

	// prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	id := request.GetFirst("id", nil)

	exists, _ := users.Exists(id)
	if exists {
		users.Delete(id)
	} else {
		message = "Record not found!"
		status = http.StatusNotFound
	}

	// prepare the data
	if request.IsJson() && request.WantsJson() {
		response.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
	}

	request.SetFlash("success", message)
	request.RedirectPrevious()
}
