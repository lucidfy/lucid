package users_handler

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/app/validations"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
)

func create(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := engine.Session
	req := engine.Request
	res := engine.Response

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

	if req.WantsJson() {
		return res.Json(data, http.StatusOK)
	}

	return res.View(
		[]string{"base", "users/show"},
		data,
	)
}

func store(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.ResponseWriter
	// r := engine.HttpRequest
	ses := engine.Session
	req := engine.Request
	res := engine.Response
	url := engine.URL

	//> validate the inputs
	if validator := req.Validator(validations.Users().Create()); validator != nil {
		if req.WantsJson() {
			return res.Json(map[string]interface{}{
				"fails": validator.ValidationError,
			}, http.StatusUnauthorized)
		}

		ses.PutFlashMap("fails", validator.ValidationError)
		ses.PutFlashMap("inputs", req.All())
		return url.RedirectPrevious()
	}

	//> prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	//> create user based on all request inputs
	data, app_err := users.Create(req.All())
	if app_err != nil {
		return app_err
	}

	//> for api based
	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    data,
		}, status)
	}

	//> for form based, just redirect
	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
