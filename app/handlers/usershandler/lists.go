package usershandler

import (
	"net/http"
	"strconv"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
	"github.com/gorilla/csrf"
)

const PER_PAGE = 2

func Lists(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	req := request.Parse(r)

	currentPage, err := strconv.Atoi(req.GetFirst("page", "1"))
	if errors.Handler("getting query 'page' error", err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	records, err := users.Lists(currentPage, PER_PAGE, "id", "desc")
	if errors.Handler("cannot fetch users list", err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usersTable, err := response.HTML([]string{"users/_table.go.html"}, map[string]interface{}{
		"records":        records,
		csrf.TemplateTag: csrf.TemplateField(r),
	})
	if errors.Handler("users table not working", err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// prepare the data
	data := map[string]interface{}{
		"title":      "Users List",
		"records":    records,
		"usersTable": usersTable,
		"pagination": records.Links(),
	}

	// this is api request
	if req.IsJson() && req.WantsJson() {
		response.Json(w, data, http.StatusOK)
		return
	}

	response.View(
		w,
		[]string{
			"base.go.html",
			"users/lists.go.html",
			// OTHER WAY of parsing the component
			// "users/_table.go.html",
		},
		data,
	)
}