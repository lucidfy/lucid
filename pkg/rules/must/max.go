package must

import (
	"fmt"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type Max struct {
	CustomErrorMessage func(string, string, int) string
	Value              int
	Translation        *lang.Translations
}

func (r *Max) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.Value)
	}
	return r.Translation.Get("validations.max", helpers.MS{
		":field":  inputField,
		":value":  inputValue,
		":length": fmt.Sprint(r.Value),
	})
}

func (r *Max) Valid(inputField string, inputValue string) bool {
	return len(inputValue) <= r.Value
}

func (r *Max) SetTranslation(t *lang.Translations) {
	r.Translation = t
}
