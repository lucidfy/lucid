package sample_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/pkg/lucid"
	"github.com/lucidfy/lucid/pkg/rules/must"
)

var RequestRoute = routes.Routing{
	Path:    "/samples/requests",
	Name:    "",
	Method:  routes.Method{"GET", "POST"},
	Handler: sampleRequests,
}

func sampleRequests(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	req := engine.GetRequest()
	res := engine.GetResponse()

	// prepare the data
	data := helpers.MP{
		"req.All()":         req.All(),
		"req.Get(k string)": req.Get("email"),
		"req.GetFirst(k string, dfault interface{})": req.GetFirst("email", nil),
		"req.Input(k string, dfault interface{})":    req.Input("email", nil),
		"req.HasContentType(substr string)":          req.HasContentType("json"),
		"req.HasAccept(substr string)":               req.HasAccept("json"),
		"req.IsForm()":                               req.IsForm(),
		"req.IsJson()":                               req.IsJson(),
		"req.IsMultipart()":                          req.IsMultipart(),
		"req.WantsJson()":                            req.WantsJson(),
		"req.Validator(setOfRules *must.SetOfRules)": req.Validator(&must.SetOfRules{
			"name": {
				&must.Required{},
				&must.Min{Value: 4},
			},
		}),
		"req.GetIp()":        req.GetIp(),
		"req.GetUserAgent()": req.GetUserAgent(),
	}

	return res.Json(data, http.StatusOK)
}
