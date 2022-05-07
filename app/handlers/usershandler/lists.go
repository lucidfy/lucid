package usershandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/session"
	"github.com/lucidfy/lucid/pkg/facade/urls"
	"github.com/lucidfy/lucid/pkg/paginate/searchable"
)

const PAGE = "1"
const PER_PAGE = "5"
const SORT_COLUMN = "id"
const SORT_TYPE = "desc"

func Lists(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	//> prepare the searchable structure
	searchable, err := prepare(req, url)
	if errors.Handler("error preparing searchable table", err) {
		return &errors.AppError{Error: err, Message: "error preparing searchable table", Code: http.StatusInternalServerError}
	}

	//> fetch the searchable
	err = users.Lists(searchable)
	if errors.Handler("error fetching users list", err) {
		return &errors.AppError{Error: err, Message: "error fetching users list", Code: http.StatusInternalServerError}
	}

	data := map[string]interface{}{
		"title":          "Users List",
		"data":           searchable,
		"links_array":    searchable.Paginate.ToArray(),
		"success":        ses.GetFlash("success"),
		"error":          ses.GetFlash("error"),
		csrf.TemplateTag: csrf.TemplateField(r),
	}

	//> here, we determine if the requestor wants a json response
	if req.WantsJson() {
		return res.Json(data, http.StatusOK)
	}

	//> or else, provide an html response instead.
	return res.View(
		[]string{
			"base",
			"users/lists",
			"users/_table",
		},
		data,
	)
}

func prepare(request request.MuxRequest, url urls.MuxUrl) (*searchable.Table, error) {
	//> get the current "page", literally the default of each current page should always be 1
	currentPage, err := strconv.Atoi(request.Input("page", PAGE).(string))
	if err != nil {
		return nil, err
	}

	//> get the "per-page", though the default will be relying to defaultPerPage
	//> then check if the per page reaches the cap of 20 records per page
	//> if ever someone tries to bypass the value, we over-write it to 20
	perPage, err := strconv.Atoi(request.Input("per-page", PER_PAGE).(string))
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
	st.Paginate.BaseUrl = request.Input("pagination_url", url.FullUrl()).(string)

	orderByCol := request.Input("sort-column", SORT_COLUMN).(string)
	orderBySort := request.Input("sort-type", SORT_TYPE).(string)
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
				Value:         request.Input("search[name]", "").(string),
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
				Value:         request.Input("search[email]", "").(string),
				CanSearch:     true,
				SearchColumn:  []string{"email"},
				SearchPattern: "-",
			},
		},
		{
			Name: "created_at",
			Input: searchable.Input{
				Visible:     true,
				Placeholder: "Created At",
				CanSearch:   false,
			},
		},
		{
			Name: "updated_at",
			Input: searchable.Input{
				Visible:     true,
				Placeholder: "Updated At",
				CanSearch:   false,
			},
		},
		{
			Name: "search",
			Input: searchable.Input{
				Visible:       false,
				Placeholder:   "*Search*",
				Value:         request.Input("search", "").(string),
				CanSearch:     true,
				SearchColumn:  []string{"email", "name"},
				SearchPattern: "<->",
			},
		},
		{
			Name: "page",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("page", PAGE).(string),
				CanSearch: false,
			},
		},
		{
			Name: "per-page",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("per-page", PER_PAGE).(string),
				CanSearch: false,
			},
		},
		{
			Name: "sort-column",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("sort-column", SORT_COLUMN).(string),
				CanSearch: false,
			},
		},
		{
			Name: "sort-type",
			Input: searchable.Input{
				Visible:   false,
				Value:     request.Input("sort-type", SORT_TYPE).(string),
				CanSearch: false,
			},
		},
	}
	return st
}
