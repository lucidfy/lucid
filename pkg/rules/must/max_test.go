package must

import (
	"fmt"
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
)

func init() {
	env.LoadEnv()
}

func TestMaxValid(t *testing.T) {
	rule := Max{Value: 10}

	if !rule.Valid("input_name", "a") {
		t.Errorf(`Should be valid since the letter is only "a" is at least 10`)
	}

	if !rule.Valid("input_name", "0123456789") {
		t.Errorf(`Should be valid since the number is at equal to 10 length`)
	}
}

func TestMaxInvalid(t *testing.T) {
	rule := Max{Value: 10}

	if rule.Valid("input_name", "Nisi ad ex labore reprehenderit.") {
		t.Errorf(`Should be invalid since the letter is beyond 10 length`)
	}

	if rule.Valid("input_name", "01234567890a") {
		t.Errorf(`Should be invalid since the letter is beyond 10 length`)
	}
}

func TestMaxWithDefaultErrorMessage(t *testing.T) {
	rule := Max{Value: 10}

	wantedErrMsg := "input_name is set to maximum of 10 length!"
	gotErrMsg := rule.ErrorMessage("input_name", "1234567890a")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}

func TestMaxWithCustomErrorMessage(t *testing.T) {
	rule := Max{
		Value: 10,
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
