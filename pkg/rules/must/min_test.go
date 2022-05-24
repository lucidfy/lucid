package must

import (
	"fmt"
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/facade/lang"
)

func init() {
	env.LoadEnv()
}

func TestMinValid(t *testing.T) {
	rule := Min{Value: 5, Translation: lang.Load(TestLanguages)}

	if !rule.Valid("input_name", "12345") {
		t.Errorf(`Should be valid`)
	}

	if !rule.Valid("input_name", "Nisi ad ex labore reprehenderit.") {
		t.Errorf(`Should be valid`)
	}
}

func TestMinInvalid(t *testing.T) {
	rule := Min{Value: 5, Translation: lang.Load(TestLanguages)}

	if rule.Valid("input_name", "Nis4") {
		t.Errorf(`Should be invalid`)
	}

	if rule.Valid("input_name", "1") {
		t.Errorf(`Should be invalid`)
	}

	if rule.Valid("input_name", "") {
		t.Errorf(`Should be invalid`)
	}
}

func TestMinWithDefaultErrorMessage(t *testing.T) {
	rule := Min{Value: 5, Translation: lang.Load(TestLanguages)}

	wantedErrMsg := "input_name is set to minimum of 5 length!"
	gotErrMsg := rule.ErrorMessage("input_name", "1234567890a")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}

func TestMinWithCustomErrorMessage(t *testing.T) {
	rule := Min{
		Translation: lang.Load(TestLanguages),
		Value:       10,
		CustomErrorMessage: func(field string, value string, length int) string {
			return fmt.Sprintf("This %s field is invalid!!! with value %s", field, value)
		},
	}

	wantedErrMsg := "This input_name field is invalid!!! with value 1234567890a"
	gotErrMsg := rule.ErrorMessage("input_name", "1234567890a")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}
