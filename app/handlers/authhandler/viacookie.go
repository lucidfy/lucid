package authhandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/hash"
	"github.com/daison12006013/gorvel/pkg/helpers"
)

func ViaCookie(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	res := engine.Response

	email := req.Input("email", nil).(string)
	password := req.Input("password", nil).(string)

	if appErr := users.Exists("email", &email); appErr != nil {
		appErr.Message = "Email or Password is incorrect!"
		return appErr
	}

	data, appErr := users.Find(&email, "email")
	if appErr != nil {
		appErr.Message = "Email or Password is incorrect!"
		return appErr
	}

	record := data.Model

	helpers.DD(
		"user found:",
		record,
		hash.Check(password, record.Password),
	)

	// data := helpers.MP{
	// 	"req.All()":          req.All(),
	// 	"req.IsForm()":       req.IsForm(),
	// 	"req.IsJson()":       req.IsJson(),
	// 	"req.IsMultipart()":  req.IsMultipart(),
	// 	"req.WantsJson()":    req.WantsJson(),
	// 	"req.GetIp()":        req.GetIp(),
	// 	"req.GetUserAgent()": req.GetUserAgent(),
	// }

	return res.Json(data, http.StatusOK)
}
