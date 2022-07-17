package cache

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnvForTests()
}

func TestFileCache(t *testing.T) {
	store := Store("file")
	store.Put("hello", "world")
	store.Put("users_online", map[string]interface{}{
		"online": 1000,
		"female": 500,
		"male":   500,
	})

	got, app_err := store.Get("hello")
	expect := "world"
	if app_err != nil {
		t.Error("Error getting cache Get")
	}
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}

	got, _ = store.Get("users_online")
	expect = `{"female":500,"male":500,"online":1000}`
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}

	m := &map[string]interface{}{}
	store.GetAs("users_online", m)

	expects := map[string]float64{
		"online": 1000,
		"female": 500,
		"male":   500,
	}
	for k, v := range expects {
		if (*m)[k] != v {
			t.Errorf("got %q, expect %f", (*m)[k], v)
		}
	}
}
