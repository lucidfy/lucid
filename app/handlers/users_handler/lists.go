package users_handler

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/pkg/paginate/searchable"
)

const PAGE = "1"
const PER_PAGE = "5"
const SORT_COLUMN = "id"
const SORT_TYPE = "desc"

func lists(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := engine.Session
	req := engine.Request
	res := engine.Response

	//> prepare the searchable structure
	searchable, app_err := prepare(T)
	if app_err != nil {
		return app_err
	}

	//> fetch the searchable
	app_err = users.Lists(searchable)
	if app_err != nil {
		return app_err
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

func prepare(T engines.EngineContract) (*searchable.Table, *errors.AppError) {
	engine := T.(engines.NetHttpEngine)
	req := engine.Request
	url := engine.URL

	//> get the current "page", literally the default of each current page should always be 1
	current_page, app_err := helpers.StringToInt(req.Input("page", PAGE).(string))
	if app_err != nil {
		return nil, app_err
	}

	//> get the "per-page", though the default will be relying to defaultPerPage
	//> then check if the per page reaches the cap of 20 records per page
	//> if ever someone tries to bypass the value, we over-write it to 20
	per_page, app_err := helpers.StringToInt(req.Input("per-page", PER_PAGE).(string))

	if app_err != nil {
		return nil, app_err
	}

	if per_page > 20 {
		per_page = 20
	}

	var st searchable.Table
	st = *table(T, &st)

	st.Paginate.CurrentPage = current_page
	st.Paginate.PerPage = per_page
	st.Paginate.BaseURL = req.Input("pagination_url", url.FullURL()).(string)

	order_by_col := req.Input("sort-column", SORT_COLUMN).(string)
	st.OrderByCol = &order_by_col

	order_by_sort := req.Input("sort-type", SORT_TYPE).(string)
	st.OrderBySort = &order_by_sort

	return &st, nil
}

func table(T engines.EngineContract, st *searchable.Table) *searchable.Table {
	engine := T.(engines.NetHttpEngine)
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
