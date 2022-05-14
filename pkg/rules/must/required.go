package must

import (
	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

type Required struct {
	CustomErrorMessage func(string, string) string
}

func (r *Required) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}
	return translations.T("validations.required", helpers.MS{
		":field": inputField,
		":value": inputValue,
	})
}

func (r *Required) Valid(inputField string, inputValue string) bool {
	return len(inputValue) > 0
}
