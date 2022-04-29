package hash

import (
	"testing"

	"github.com/daison12006013/lucid/pkg/env"
)

func init() {
	env.LoadEnv()
}

func TestHash(t *testing.T) {
	password := "1234qwer"

	hashed, err := Make(password)

	if err != nil {
		t.Errorf("hash.Make is not working %q", err)
	}

	if Check(password, hashed) == false {
		t.Errorf("hash.Check is not working")
	}
}
