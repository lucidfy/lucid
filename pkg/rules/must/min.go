package must

import (
	"fmt"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type Min struct {
	CustomErrorMessage func(string, string, int) string
	Value              int
	Translation        *lang.Translations
}

func (r *Min) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.Value)
	}
	return r.Translation.Get("validations.min", helpers.MS{
		":field":  inputField,
		":value":  inputValue,
		":length": fmt.Sprint(r.Value),
	})
}

func (r *Min) Valid(inputField string, inputValue string) bool {
	return len(inputValue) >= r.Value
}

func (r *Min) SetTranslation(t *lang.Translations) {
	r.Translation = t
}
