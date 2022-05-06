package validations

import "github.com/lucidfy/lucid/pkg/rules/must"

func AuthValidateLogin() *must.SetOfRules {
	return &must.SetOfRules{
		"email": {
			&must.Email{},
		},
		"password": {
			&must.Min{Value: 5},
		},
	}
}
