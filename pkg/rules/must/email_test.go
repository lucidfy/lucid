package must

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	rule := Email{}

	if rule.Valid("email", "johndoe") == true {
		t.Errorf(`Email should be invalid as we used "johndoe"`)
	}
}

func TestEmailWithAtSign(t *testing.T) {
	rule := Email{}

	if rule.Valid("email", "johndoe@") == true {
		t.Errorf(`Email should be invalid as we used "johndoe@"`)
	}
}

func TestEmailWithUserDomain(t *testing.T) {
	rule := Email{}

	if rule.Valid("email", "johndoe@domain") {
		t.Errorf(`Email should be invalid as we used "johndoe@domain"`)
	}
}

func TestEmailWithFullDomain(t *testing.T) {
	rule := Email{}

	if !rule.Valid("email", "johndoe@domain.com") {
		t.Errorf(`Email should be valid as we used "johndoe@domain.com"`)
	}
}

func TestEmailWithPlusAndFullDomain(t *testing.T) {
	rule := Email{}

	if !rule.Valid("email", "johndoe+test1@domain.com") {
		t.Errorf(`Email should be valid as we used "johndoe+test1@domain.com"`)
	}
}

func TestEmailWithCustomErrorMessage(t *testing.T) {
	rule := Email{
		CustomErrorMessage: func(field string, value string) string {
			return fmt.Sprintf("This %s field is invalid!!! with value %s", field, value)
		},
	}

	if rule.Valid("email", "johndoe+ @ email . com") {
		t.Errorf(`Email should be invalid as we used "johndoe+ @ email . com"`)
	}

	wantedErrMsg := "This email field is invalid!!! with value johndoe+ @ email . com"
	gotErrMsg := rule.ErrorMessage("email", "johndoe+ @ email . com")
	if gotErrMsg != wantedErrMsg {
		t.Errorf(`got %q, wanted %q`, gotErrMsg, wantedErrMsg)
	}
}
