package validations

import (
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type UsersValidator struct {
	Rules *must.SetOfRules
}

func Users() *UsersValidator {
	return &UsersValidator{
		Rules: &must.SetOfRules{
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
		},
	}
}

func (v UsersValidator) Create() *must.SetOfRules {
	return v.Rules
}

func (v UsersValidator) Update() *must.SetOfRules {
	sets := *v.Rules

	sets["password"] = []must.Rule{
		&must.Min{Value: 5},
		&must.StrictPassword{
			WithSpecialChar: true,
			WithUpperCase:   true,
			WithLowerCase:   true,
			WithDigit:       true,
		},
	}

	return &sets
}
