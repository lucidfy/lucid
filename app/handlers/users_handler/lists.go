package users_handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/session"
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

	//> prepare the searchable structure
	searchable, err := prepare(T)
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

func prepare(T engines.EngineContract) (*searchable.Table, error) {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	url := engine.Url

	//> get the current "page", literally the default of each current page should always be 1
	currentPage, err := strconv.Atoi(req.Input("page", PAGE).(string))
	if err != nil {
		return nil, err
	}

	//> get the "per-page", though the default will be relying to defaultPerPage
	//> then check if the per page reaches the cap of 20 records per page
	//> if ever someone tries to bypass the value, we over-write it to 20
	perPage, err := strconv.Atoi(req.Input("per-page", PER_PAGE).(string))
	if err != nil {
		return nil, err
	}
	if perPage > 20 {
		perPage = 20
	}

	var st searchable.Table
	st = *table(T, &st)

	st.Paginate.CurrentPage = currentPage
	st.Paginate.PerPage = perPage
	st.Paginate.BaseUrl = req.Input("pagination_url", url.FullUrl()).(string)

	orderByCol := req.Input("sort-column", SORT_COLUMN).(string)
	orderBySort := req.Input("sort-type", SORT_TYPE).(string)
	st.OrderByCol = &orderByCol
	st.OrderBySort = &orderBySort

	return &st, nil
}

func table(T engines.EngineContract, st *searchable.Table) *searchable.Table {
	engine := T.(engines.MuxEngine)
	req := engine.Request

	st.Headers = []searchable.Header{
		{
			Name: "name",
			Input: searchable.Input{
				Visible:       true,
				Placeholder:   "Name*",
				Value:         req.Input("search[name]", "").(string),
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
				Value:         req.Input("search[email]", "").(string),
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
				Value:         req.Input("search", "").(string),
				CanSearch:     true,
				SearchColumn:  []string{"email", "name"},
				SearchPattern: "<->",
			},
		},
		{
			Name: "page",
			Input: searchable.Input{
				Visible:   false,
				Value:     req.Input("page", PAGE).(string),
				CanSearch: false,
			},
		},
		{
			Name: "per-page",
			Input: searchable.Input{
				Visible:   false,
				Value:     req.Input("per-page", PER_PAGE).(string),
				CanSearch: false,
			},
		},
		{
			Name: "sort-column",
			Input: searchable.Input{
				Visible:   false,
				Value:     req.Input("sort-column", SORT_COLUMN).(string),
				CanSearch: false,
			},
		},
		{
			Name: "sort-type",
			Input: searchable.Input{
				Visible:   false,
				Value:     req.Input("sort-type", SORT_TYPE).(string),
				CanSearch: false,
			},
		},
	}
	return st
}
