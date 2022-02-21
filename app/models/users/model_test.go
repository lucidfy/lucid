package users

import (
	"testing"

	"github.com/daison12006013/gorvel/pkg/env"
)

func init() {
	env.LoadEnv()
}

func TestLists(t *testing.T) {
	id := "id"
	desc := "desc"

	var paginated Paginate
	paginated.CurrentPage = 1
	paginated.BaseUrl = ""
	paginated.PerPage = 3
	paginated.OrderByCol = &id
	paginated.OrderBySort = &desc
	err := Lists(&paginated)

	if err != nil {
		t.Errorf("paginated lists is not working %q", err)
	}

	if paginated.Total == 0 {
		t.Errorf("paginated.Total should not be Zero")
	}

	if paginated.PerPage == 0 {
		t.Errorf("paginated.PerPage should not be Zero")
	}

	if paginated.CurrentPage == 0 {
		t.Errorf("paginated.CurrentPage should not be Zero")
	}

	if paginated.LastPage == 0 {
		t.Errorf("paginated.LastPage should not be Zero")
	}

	if len(*paginated.Items.(*[]Attributes)) == 0 {
		t.Errorf("paginated.Items should have a record")
	}
}
