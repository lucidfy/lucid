package validations

import (
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type ReportsValidator struct {
	Rules *must.SetOfRules
}

func Reports() *ReportsValidator {
	return &ReportsValidator{
		Rules: &must.SetOfRules{
			// "form_input": {
			// 	&must.Required{},
			// },
		},
	}
}

func (v ReportsValidator) Create() *must.SetOfRules {
	return v.Rules
}

func (v ReportsValidator) Update() *must.SetOfRules {
	sets := *v.Rules

	// sets["form_input"] = []must.Rule{}

	return &sets
}
