package usershandler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

const PER_PAGE = 2

func Lists(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	req := request.Parse(r)

	currentPage, err := strconv.Atoi(req.GetFirst("page"))
	if errHandle(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	records, err := users.Lists(currentPage, PER_PAGE, "id", "desc")
	if errHandle(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// prepare the data
	data := map[string]interface{}{
		"title":   "Lists",
		"records": records,
	}

	// this is api request
	if req.IsJson() && req.WantsJson() {
		response.Json(w, data, http.StatusOK)
		return
	}

	response.View(
		w,
		[]string{"base.go.tmpl", "users/lists.go.tmpl"},
		data,
	)
}

func errHandle(err error) bool {
	if err != nil {
		// if we're on debugging mode, just throw the error
		if os.Getenv("APP_DEBUG") == "true" {
			logger.Fatal(err)
		}
		return true
	}
	return false
}
