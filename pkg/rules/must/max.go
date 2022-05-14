package must

import (
	"fmt"

	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

type Max struct {
	CustomErrorMessage func(string, string, int) string
	Value              int
}

func (r *Max) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.Value)
	}
	return translations.T("validations.max", helpers.MS{
		":field":  inputField,
		":value":  inputValue,
		":length": fmt.Sprint(r.Value),
	})
}

func (r *Max) Valid(inputField string, inputValue string) bool {
	return len(inputValue) <= r.Value
}
