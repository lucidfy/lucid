package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/app/validations"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/gorilla/csrf"
)

func Create(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	r := engine.HttpRequest
	// req := engine.Request
	res := engine.Response
	ses := engine.Session

	data := map[string]interface{}{
		"title":          "Create Form",
		"record":         &users.Model{},
		"isCreate":       true,
		csrf.TemplateTag: csrf.TemplateField(r),

		//> to retrieve the flashes from Store() or somewhere
		//> that redirects back to Create()
		"success": ses.GetFlash("success"),
		"error":   ses.GetFlash("error"),
		"fails":   ses.GetFlashMap("fails"),
		"inputs":  ses.GetFlashMap("inputs"),
	}

	return res.View(
		[]string{"base", "users/show"},
		data,
	)
}

func Store(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response
	ses := engine.Session
	url := engine.Url

	//> validate the inputs
	validator := req.Validator(validations.UserValidateCreate())
	if validator != nil {
		ses.SetFlashMap("fails", validator.ValidationError)
		ses.SetFlashMap("inputs", req.All())
		url.RedirectPrevious()
		return nil
	}

	//> prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	//> create user based on all request inputs
	data, appErr := users.Create(req.All())
	if appErr != nil {
		return appErr
	}

	//> for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
			"data":    data,
		}, status)
	}

	//> for form based, just redirect
	ses.SetFlash("success", message)
	url.RedirectPrevious()
	return nil
}
