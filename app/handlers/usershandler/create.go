package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
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
		"success":        req.GetFlash("success"),
		"error":          req.GetFlash("error"),
		csrf.TemplateTag: csrf.TemplateField(r),
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
