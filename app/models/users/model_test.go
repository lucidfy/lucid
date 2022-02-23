package users

import (
	"testing"

	"github.com/daison12006013/gorvel/pkg/env"
	"github.com/daison12006013/gorvel/pkg/paginate/searchable"
)

func init() {
	env.LoadEnv()
}

func TestLists(t *testing.T) {
	id := "id"
	desc := "desc"

	var s searchable.Table
	s.Paginate.CurrentPage = 1
	s.Paginate.BaseUrl = ""
	s.Paginate.PerPage = 3
	s.OrderByCol = &id
	s.OrderBySort = &desc
	err := Lists(&s)

	if err != nil {
		t.Errorf("paginated lists is not working %q", err)
	}

	if s.Paginate.Total == 0 {
		t.Errorf("s.Paginate.Total should not be Zero")
	}

	if s.Paginate.PerPage == 0 {
		t.Errorf("s.Paginate.PerPage should not be Zero")
	}

	if s.Paginate.CurrentPage == 0 {
		t.Errorf("s.Paginate.CurrentPage should not be Zero")
	}

	if s.Paginate.LastPage == 0 {
		t.Errorf("s.Paginate.LastPage should not be Zero")
	}

	if len(*s.Paginate.Items.(*[]Model)) == 0 {
		t.Errorf("s.Paginate.Items should have a record")
	}
}
