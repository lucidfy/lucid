package validations

import "github.com/lucidfy/lucid/pkg/rules/must"

type AuthValidator struct {
	Rules *must.SetOfRules
}

func Auth() *AuthValidator {
	return &AuthValidator{
		Rules: &must.SetOfRules{
			"email": {
				&must.Required{},
				&must.Email{},
			},
			"password": {
				&must.Required{},
				&must.Min{Value: 5},
			},
		},
	}
}

// Login, by default we will only require the email
// and password
func (v AuthValidator) Login() *must.SetOfRules {
	return v.Rules
}

// SignUp, when signing up, we must append
// the "confirm_password" and the rule must
// match with field "password"
func (v AuthValidator) SignUp() *must.SetOfRules {
	sets := *v.Rules

	sets["confirm_password"] = []must.Rule{
		&must.Required{},
		&must.Matches{TargetField: "password"},
	}

	return &sets
}
