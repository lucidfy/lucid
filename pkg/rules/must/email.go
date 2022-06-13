package must

import (
	"regexp"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type Email struct {
	CustomErrorMessage func(string, string) string
	Translation        *lang.Translations
}

func (r *Email) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}
	return r.Translation.Get("validations.email", helpers.MS{
		":field": inputField,
		":value": inputValue,
	})
}

func (r *Email) Valid(inputField string, inputValue string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(inputValue)
}

func (r *Email) SetTranslation(t *lang.Translations) {
	r.Translation = t
}
