package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
)

func Delete(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	request := engine.Request
	response := engine.Response

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
