package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/app/validations"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/gorilla/csrf"
)

func Create(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	data := map[string]interface{}{
		"title":          "Create Form",
		"record":         &users.Model{},
		"isCreate":       true,
		csrf.TemplateTag: csrf.TemplateField(r),

		// to retrieve the flashes from Store() or somewhere
		// that redirects back to Create()
		"success": req.GetFlash("success"),
		"error":   req.GetFlash("error"),
		"fails":   req.GetFlashMap("fails"),
		"inputs":  req.GetFlashMap("inputs"),
	}

	return res.View(
		[]string{"base", "users/show"},
		data,
	)
}

func Store(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// validate the inputs
	validator := req.Validator(validations.UserValidateCreate())
	if validator != nil {
		req.SetFlashMap("fails", validator.ValidationError)
		req.SetFlashMap("inputs", req.All())
		req.RedirectPrevious()
		return nil
	}

	message := "Successfully Created!"
	status := http.StatusOK

	data, appErr := users.Create(req.All())
	if appErr != nil {
		return appErr
	}

	// for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
			"data":    data,
		}, status)
	}

	// for form based, just redirect
	req.SetFlash("success", message)
	req.RedirectPrevious()
	return nil
}
