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

const PAGE = "1"
const PER_PAGE = "5"
const SORT_COLUMN = "id"
const SORT_TYPE = "desc"

func Lists(w http.ResponseWriter, r *http.Request) {
	// let's extend the request
	rp := request.Parse(r)

	// prepare the searchable structure
	searchable, e := prepare(rp)
	if errors.Handler("getting query 'page' error", e) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// fetch the searchable
	err := users.Lists(searchable)
	if errors.Handler("cannot fetch users list", err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// here, we determine if the requestor wants a json response
	if rp.IsJson() && rp.WantsJson() {
		response.Json(w, searchable.Paginate.ToArray(), http.StatusOK)
		return
	}

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
			"data":           searchable,
			csrf.TemplateTag: csrf.TemplateField(r),
		},
	)
}

// --------
//
func prepare(rp request.ParsedRequest) (*users.SearchableTable, error) {
	// get the current "page", literally the default of each current page should always be 1
	currentPage, err := strconv.Atoi(rp.Input("page", PAGE))
	if err != nil {
		return nil, err
	}

	// get the "per-page", though the default will be relying to defaultPerPage
	// then check if the per page reaches the cap of 20 records per page
	// if ever someone tries to bypass the value, we over-write it to 20
	perPage, err := strconv.Atoi(rp.Input("per-page", PER_PAGE))
	if err != nil {
		return nil, err
	}
	if perPage > 20 {
		perPage = 20
	}

	var s users.SearchableTable
	s = *data(rp, &s)

	s.Paginate.CurrentPage = currentPage
	s.Paginate.PerPage = perPage
	s.Paginate.BaseUrl = rp.FullUrl()

	orderByCol := rp.Input("sort-column", SORT_COLUMN)
	orderBySort := rp.Input("sort-type", SORT_TYPE)
	s.OrderByCol = &orderByCol
	s.OrderBySort = &orderBySort

	return &s, nil
}

// --------
//
func data(rp request.ParsedRequest, searchable *users.SearchableTable) *users.SearchableTable {
	searchable.Headers = []users.Header{
		{
			Name: "name",
			Input: users.Input{
				Visible:       true,
				Placeholder:   "Name*",
				Value:         rp.Input("search[name]", ""),
				CanSearch:     true,
				SearchColumn:  []string{"name"},
				SearchPattern: "->",
			},
		},
		{
			Name: "email",
			Input: users.Input{
				Visible:       true,
				Placeholder:   "Email",
				Value:         rp.Input("search[email]", ""),
				CanSearch:     true,
				SearchColumn:  []string{"email"},
				SearchPattern: "-",
			},
		},
		{
			Name: "search",
			Input: users.Input{
				Visible:       false,
				Placeholder:   "*Search*",
				Value:         rp.Input("search", ""),
				CanSearch:     true,
				SearchColumn:  []string{"email", "name"},
				SearchPattern: "<->",
			},
		},
		{
			Name: "page",
			Input: users.Input{
				Visible:   false,
				Value:     rp.Input("page", PAGE),
				CanSearch: false,
			},
		},
		{
			Name: "per-page",
			Input: users.Input{
				Visible:   false,
				Value:     rp.Input("per-page", PER_PAGE),
				CanSearch: false,
			},
		},
		{
			Name: "sort-column",
			Input: users.Input{
				Visible:   false,
				Value:     rp.Input("sort-column", SORT_COLUMN),
				CanSearch: false,
			},
		},
		{
			Name: "sort-type",
			Input: users.Input{
				Visible:   false,
				Value:     rp.Input("sort-type", SORT_TYPE),
				CanSearch: false,
			},
		},
	}
	return searchable
}
