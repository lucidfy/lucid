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

var defaultCurrentPage string = "1"
var defaultPerPage string = "5"
var defaultSortCol string = "id"
var defaultSortType string = "desc"

func Lists(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	rp := request.Parse(r)

	// prepare the paginated structure
	records, e := prepare(rp)
	if errors.Handler("getting query 'page' error", e) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// fetch the records
	err := users.Lists(records)
	if errors.Handler("cannot fetch users list", err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// here, we determine if the requestor wants a json response
	if rp.IsJson() && rp.WantsJson() {
		response.Json(w, records.ToArray(), http.StatusOK)
		return
	}

	/*
		*ALTERNATIVE: way of rendering the user's table
		tables, err := response.HTML([]string{"users/_table.go.html"}, map[string]interface{}{
			"records":        records,
			csrf.TemplateTag: csrf.TemplateField(r),
		})
		if errors.Handler("users table not working", err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	*/

	// or else, provide an html response instead.
	response.View(
		w,
		[]string{
			"base.go.html",
			"users/lists.go.html",
			"users/_table.go.html",
		},
		map[string]interface{}{
			"title":          "Users List",
			"records":        records,
			csrf.TemplateTag: csrf.TemplateField(r),

			// *ALTERNATIVE
			// "tables": tables,
		},
	)
}

func prepare(rp request.ParsedRequest) (*users.Paginate, error) {
	// get the current "page", literally the default of each current page should always be 1
	currentPage, err := strconv.Atoi(*rp.GetFirst("page", &defaultCurrentPage))
	if err != nil {
		return nil, err
	}

	// get the "per-page", though the default will be relying to defaultPerPage
	// then check if the per page reaches the cap of 20 records per page
	// if ever someone tries to bypass the value, we over-write it to 20
	perPage, err := strconv.Atoi(*rp.GetFirst("per-page", &defaultPerPage))
	if err != nil {
		return nil, err
	}
	if perPage > 20 {
		perPage = 20
	}

	orderByCol := *rp.GetFirst("sort-column", &defaultSortCol)
	orderBySort := *rp.GetFirst("sort-type", &defaultSortType)

	var p users.Paginate
	p.CurrentPage = currentPage
	p.PerPage = perPage
	p.OrderByCol = &orderByCol
	p.OrderBySort = &orderBySort
	p.TextSearch = rp.GetFirst("search", nil)
	p.BaseUrl = rp.FullUrl()

	return &p, nil
}
