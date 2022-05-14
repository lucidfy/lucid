package rules

import (
	"testing"

	"github.com/lucidfy/lucid/pkg/rules/must"
)

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
		t.Error("Validation should pass")
	}

	// ---

	inputValues = map[string]interface{}{
		"name":             "",
		"email":            "johndoe",
		"password":         "1234qwerASDF!@#$",
		"confirm_password": "1q3rZo4ogF!t4$",
	}

	validationErrors = GetErrors(setOfRules, inputValues)

	if validationErrors["confirm_password"] != "confirm_password did not match with password" {
		t.Error("Validating confirm_password seems not right!")
	}

	if validationErrors["email"] != "johndoe is not a valid email address!" {
		t.Error("Validating email seems not right!")
	}

	if validationErrors["name"] != "name is set to minimum of 4 length" {
		t.Error("Validating email seems not right!")
	}
}
