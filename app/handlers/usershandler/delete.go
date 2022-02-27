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

	id := request.Input("id", nil).(string)

	exists, _ := users.Exists(&id)
	if exists {
		users.Find(&id).Delete()
	} else {
		message = "Record not found!"
		status = http.StatusNotFound
	}

	// for api based
	if request.IsJson() && request.WantsJson() {
		response.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
		return
	}

	// for form based, just redirect
	request.SetFlash("success", message)
	request.RedirectPrevious()
}
