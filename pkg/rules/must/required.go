package must

import (
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type Required struct {
	CustomErrorMessage func(string, string) string
	Translation        *lang.Translations
}

func (r *Required) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}
	return r.Translation.Get("validations.required", helpers.MS{
		":field": inputField,
		":value": inputValue,
	})
}

func (r *Required) Valid(inputField string, inputValue string) bool {
	return len(inputValue) > 0
}

func (r *Required) SetTranslation(t *lang.Translations) {
	r.Translation = t
}
