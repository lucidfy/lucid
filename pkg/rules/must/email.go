package must

import (
	"regexp"

	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

type Email struct {
	CustomErrorMessage func(string, string) string
}

func (r *Email) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}
	return translations.T("validations.email", helpers.MS{
		":field": inputField,
		":value": inputValue,
	})
}

func (r *Email) Valid(inputField string, inputValue string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(inputValue)
}
