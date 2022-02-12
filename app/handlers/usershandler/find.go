package usershandler

import (
	"net/http"
	"os"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Find(w http.ResponseWriter, r *http.Request) {
	// If we're properly writing a response from the http.ResponseWriter,
	// therefore no need to write the header as Status 200 "OK",
	// although it is still good to write it at first, and override
	// underneath if there are conditional cases that you want to filter-out
	w.WriteHeader(http.StatusOK)

	// let's extend the request
	req := request.Parse(r)

	// fetch the record in the database
	record, err := users.FindById(*req.Input("id"))
	if err != nil {
		// if we're on debugging mode, just throw the error
		if os.Getenv("APP_DEBUG") == "true" {
			logger.Fatal(err)
		}
		w.WriteHeader(http.StatusNotFound)
	}

	// prepare the data
	data := map[string]interface{}{
		"title":  record.Name + "'s Profile",
		"record": record,
	}

	// this is api request
	if req.IsJson() && req.WantsJson() {
		response.Json(w, data)
	}

	// render the template
	response.View(
		w,
		[]string{"base.html", "users/find.html"},
		data,
	)
}
