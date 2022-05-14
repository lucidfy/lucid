package must

import (
	"fmt"

	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

type Min struct {
	CustomErrorMessage func(string, string, int) string
	Value              int
}

func (r *Min) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.Value)
	}
	return translations.T("validations.min", helpers.MS{
		":field":  inputField,
		":value":  inputValue,
		":length": fmt.Sprint(r.Value),
	})
}

func (r *Min) Valid(inputField string, inputValue string) bool {
	return len(inputValue) >= r.Value
}
