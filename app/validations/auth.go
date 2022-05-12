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

func (v AuthValidator) Login() *must.SetOfRules {
	return v.Rules
}

func (v AuthValidator) SignUp() *must.SetOfRules {
	sets := *v.Rules

	sets["confirm_password"] = []must.Rule{
		&must.Required{},
	}

	return &sets
}
