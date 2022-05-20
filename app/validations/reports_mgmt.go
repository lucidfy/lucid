package validations

import (
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type ReportsMgmtValidator struct {
	Rules *must.SetOfRules
}

func ReportsMgmt() *ReportsMgmtValidator {
	return &ReportsMgmtValidator{
		Rules: &must.SetOfRules{
			// "form_input": {
			// 	&must.Required{},
			// },
		},
	}
}

func (v ReportsMgmtValidator) Create() *must.SetOfRules {
	return v.Rules
}

func (v ReportsMgmtValidator) Update() *must.SetOfRules {
	sets := *v.Rules

	// sets["form_input"] = []must.Rule{}

	return &sets
}
