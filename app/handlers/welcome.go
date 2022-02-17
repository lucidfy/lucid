package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	request := request.Parse(r)

	// prepare the data
	data := map[string]interface{}{
		"title": "Gorvel Rocks!",
	}

	// this is api request
	if request.IsJson() && request.WantsJson() {
		response.Json(w, data, http.StatusOK)
		return
	}

	// render the template
	response.View(
		w,
		// this example below, we're telling the compiler
		// to parse the base.go.tmpl first, and then parse the welcome.go.tmpl
		// therefore the defined "body" should render accordingly
		[]string{"base.go.tmpl", "welcome.go.tmpl"},
		data,
	)
}
