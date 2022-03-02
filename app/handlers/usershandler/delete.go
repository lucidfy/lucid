package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

func Delete(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	id := req.Input("id", nil).(string)
	if err := users.Exists(&id); err != nil {
		return err
	}

	users.Find(&id).Delete()

	// for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
	}

	// for form based, just redirect
	req.SetFlash("success", message)
	req.RedirectPrevious()
	return nil
}
