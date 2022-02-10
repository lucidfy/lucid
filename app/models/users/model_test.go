package users

import (
	"testing"

	"github.com/daison12006013/gorvel/internal/env"
)

func init() {
	env.LoadEnv()
}

func TestLists(t *testing.T) {
	paginated, err := Lists(1, 3, "id", "desc")

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

	if len(paginated.Data) == 0 {
		t.Errorf("paginated.Data should have a record")
	}
}
