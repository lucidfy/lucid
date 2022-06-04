package auth_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/hash"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/lucid"
)

var RouteResource = routes.Routing{
	Path: "/auth/login",
	Name: "auth-login",
	Resources: routes.Resources{
		"index": user,         //  GET    /auth/login
		"store": loginAttempt, //  POST   /auth/login
	},
}

func user(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	ses := engine.GetSession()
	res := engine.GetResponse()

	userID, app_err := ses.Get("authenticated")
	if userID == nil || app_err != nil {
		return res.Json(map[string]string{}, http.StatusOK)
	}

	data, app_err := users.Find(userID, nil)
	if app_err != nil {
		return app_err
	}

	return res.Json(map[string]interface{}{
		"user": data.Model,
	}, http.StatusOK)
}

func loginAttempt(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	email := req.Input("email", nil).(string)
	password := req.Input("password", nil).(string)

	if app_err := users.Exists("email", &email); app_err != nil {
		app_err.Message = "Email or Password is incorrect!"
		return app_err
	}

	data, app_err := users.Find(&email, "email")
	if app_err != nil {
		app_err.Message = "Email or Password is incorrect!"
		return app_err
	}

	record := data.Model

	if hash.Check(password, record.Password) {
		ses.Put("authenticated", record.ID)
	}

	message := "Successfully Logged In!"
	status := http.StatusOK

	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    record,
		}, status)
	}

	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
