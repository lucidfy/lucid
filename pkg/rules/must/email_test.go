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

func TestEmail(t *testing.T) {
	rule := Email{Translation: lang.Load(TestLanguages)}

	if rule.Valid("email", "johndoe") == true {
		t.Errorf(`Email should be invalid as we used "johndoe"`)
	}

	got := rule.ErrorMessage("email", "johndoe")
	expect := "email is not a valid email address!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestEmailWithAtSign(t *testing.T) {
	rule := Email{Translation: lang.Load(TestLanguages)}

	if rule.Valid("email", "johndoe@") == true {
		t.Errorf(`Email should be invalid as we used "johndoe@"`)
	}

	got := rule.ErrorMessage("email", "johndoe@")
	expect := "email is not a valid email address!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestEmailWithUserDomain(t *testing.T) {
	rule := Email{Translation: lang.Load(TestLanguages)}

	if rule.Valid("email", "johndoe@domain") {
		t.Errorf(`Email should be invalid as we used "johndoe@domain"`)
	}

	got := rule.ErrorMessage("email", "johndoe@domain")
	expect := "email is not a valid email address!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestEmailWithFullDomain(t *testing.T) {
	rule := Email{Translation: lang.Load(TestLanguages)}

	if !rule.Valid("email", "johndoe@domain.com") {
		t.Errorf(`Email should be valid as we used "johndoe@domain.com"`)
	}

	got := rule.ErrorMessage("email", "johndoe@domain.com")
	expect := "email is not a valid email address!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func TestEmailWithPlusAndFullDomain(t *testing.T) {
	rule := Email{Translation: lang.Load(TestLanguages)}

	if !rule.Valid("email", "johndoe+test1@domain.com") {
		t.Errorf(`Email should be valid as we used "johndoe+test1@domain.com"`)
	}
}

func TestEmailWithCustomErrorMessage(t *testing.T) {
	rule := Email{
		Translation: lang.Load(TestLanguages),
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
