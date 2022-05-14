package rules

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/rules/must"
)

func init() {
	env.LoadEnv()
}

func TestGetErrors(t *testing.T) {
	setOfRules := &must.SetOfRules{
		"name": {
			&must.Required{},
			&must.Min{Value: 4},
		},
		"email": {
			&must.Required{},
			&must.Email{},
		},
		"password": {
			&must.Required{},
			&must.Min{Value: 5},
			&must.StrictPassword{
				WithSpecialChar: true,
				WithUpperCase:   true,
				WithLowerCase:   true,
				WithDigit:       true,
			},
		},
		"confirm_password": {
			&must.Required{},
			&must.Matches{TargetField: "password"},
		},
	}

	inputValues := map[string]interface{}{
		"name":             "John Doe",
		"email":            "johndoe@email.com",
		"password":         "1234qwerASDF!@#$",
		"confirm_password": "1234qwerASDF!@#$",
	}

	validationErrors := GetErrors(setOfRules, inputValues)
	if len(validationErrors) != 0 {
		t.Error("Validation error should be empty!")
	}

	// ---

	inputValues = map[string]interface{}{
		"name":             "",
		"email":            "johndoe",
		"password":         "1234qwerASDF!@#$",
		"confirm_password": "1q3rZo4ogF!t4$",
	}
	validationErrors = GetErrors(setOfRules, inputValues)

	got := validationErrors["confirm_password"]
	expect := "confirm_password did not match with password!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}

	got = validationErrors["email"]
	expect = "email is not a valid email address!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}

	got = validationErrors["name"]
	expect = "name is set to minimum of 4 length!"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}
