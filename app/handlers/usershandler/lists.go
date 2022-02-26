package usershandler

import (
	"net/http"
	"strconv"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/paginate/searchable"
	"github.com/gorilla/csrf"
)

const PAGE = "1"
const PER_PAGE = "5"
const SORT_COLUMN = "id"
const SORT_TYPE = "desc"

func Lists(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	request := engine.Request
	response := engine.Response

	// prepare the searchable structure
	searchable, e := prepare(request)
	if errors.Handler("error preparing searchable table", e) {
		engine.HttpResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fetch the searchable
	err := users.Lists(searchable)
	if errors.Handler("error fetching users list", err) {
		engine.HttpResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	// here, we determine if the requestor wants a json response
	if request.IsJson() && request.WantsJson() {
		response.Json(searchable.Paginate.ToArray(), http.StatusOK)
		return
	}

	// or else, provide an html response instead.
	response.View(
		[]string{
			"base.go.html",
			"users/lists.go.html",
			"users/_table.go.html",
		},
		map[string]interface{}{
			"title":          "Users List",
			"data":           searchable,
			csrf.TemplateTag: csrf.TemplateField(engine.HttpRequest),
			"success":        request.GetFlash("success"),
			"error":          request.GetFlash("error"),
		},
	)
}

func prepare(request request.MuxRequest) (*searchable.Table, error) {
	// get the current "page", literally the default of each current page should always be 1
	currentPage, err := strconv.Atoi(request.Input("page", PAGE))
	if err != nil {
		return nil, err
	}

	// get the "per-page", though the default will be relying to defaultPerPage
	// then check if the per page reaches the cap of 20 records per page
	// if ever someone tries to bypass the value, we over-write it to 20
	perPage, err := strconv.Atoi(request.Input("per-page", PER_PAGE))
	if err != nil {
		return nil, err
	}
	if perPage > 20 {
		perPage = 20
	}

	var st searchable.Table
	st = *table(request, &st)

	st.Paginate.CurrentPage = currentPage
	st.Paginate.PerPage = perPage
	st.Paginate.BaseUrl = request.FullUrl()

	orderByCol := request.Input("sort-column", SORT_COLUMN)
	orderBySort := request.Input("sort-type", SORT_TYPE)
	st.OrderByCol = &orderByCol
	st.OrderBySort = &orderBySort

	return &st, nil
}

func table(request request.MuxRequest, st *searchable.Table) *searchable.Table {
	st.Headers = []searchable.Header{
		{
			Name: "name",
			Input: searchable.Input{
				Visible:       true,
				Placeholder:   "Name*",
				Value:         request.Input("search[name]", ""),
				CanSearch:     true,
				SearchColumn:  []string{"name"},
				SearchPattern: "->",
			},
		},
		{
			Name: "email",
			Input: searchable.Input{
				Visible:       true,
				Placeholder:   "Email",
				Value:         request.Input("search[email]", ""),
				CanSearch:     true,
				SearchColumn:  []string{"email"},
				SearchPattern: "-",
			},
		},
		{
			Name: "search",
			Input: searchable.Input{
				Visible:       false,
				Placeholder:   "*Search*",
				Value:         request.Input("search", ""),
				CanSearch:     true,
				SearchColumn:  []string{"email", "name"},
				SearchPattern: "<->",
			},
		},
		{
			Name: "page",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("page", PAGE),
				CanSearch: false,
			},
		},
		{
			Name: "per-page",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("per-page", PER_PAGE),
				CanSearch: false,
			},
		},
		{
			Name: "sort-column",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("sort-column", SORT_COLUMN),
				CanSearch: false,
			},
		},
		{
			Name: "sort-type",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("sort-type", SORT_TYPE),
				CanSearch: false,
			},
		},
	}
	return st
}
