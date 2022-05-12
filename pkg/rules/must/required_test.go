package must

import (
	"fmt"
	"testing"
)

func TestRequiredValid(t *testing.T) {
	rule := Required{}

	if !rule.Valid("input_name", "hello world!") {
		t.Errorf("It should be valid")
	}
}
func TestRequiredInvalid(t *testing.T) {
	rule := Required{}

	if rule.Valid("input_name", "") {
		t.Errorf("It should be invalid")
	}
}

func TestRequiredWithDefaultErrorMessage(t *testing.T) {
	rule := Required{}

	wantedErrMsg := "input_name is required!"
	gotErrMsg := rule.ErrorMessage("input_name", "")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}

func TestRequiredWithCustomErrorMessage(t *testing.T) {
	rule := Required{
		CustomErrorMessage: func(field string, value string) string {
			return fmt.Sprintf("This %s field is required!!!", field)
		},
	}

	wantedErrMsg := "This input_name field is required!!!"
	gotErrMsg := rule.ErrorMessage("input_name", "")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}
