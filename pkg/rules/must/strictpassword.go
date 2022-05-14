package must

import (
	"unicode"

	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/resources/translations"
)

type StrictPassword struct {
	Value   string
	message string

	WithSpecialChar bool
	WithUpperCase   bool
	WithLowerCase   bool
	WithDigit       bool

	ErrorMessageNoSpecialChar func(string, string) string
	ErrorMessageNoUpperCase   func(string, string) string
	ErrorMessageNoLowerCase   func(string, string) string
	ErrorMessageNoDigit       func(string, string) string
}

func (r *StrictPassword) ErrorMessage(inputField string, inputValue string) string {
	return r.message
}

func (r *StrictPassword) Valid(inputField string, inputValue string) bool {
	var hasDigit, hasUpper, hasLower, hasSpecial bool

	for _, c := range inputValue {
		switch {
		case unicode.IsNumber(c):
			hasDigit = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if r.WithSpecialChar && !hasSpecial {
		if r.ErrorMessageNoSpecialChar != nil {
			r.message = r.ErrorMessageNoSpecialChar(inputField, inputValue)
		} else {
			r.message = translations.T("validations.strictpassword.special", helpers.MS{
				":field": inputField,
				":value": inputValue,
			})

		}
		return false
	}

	if r.WithUpperCase && !hasUpper {
		if r.ErrorMessageNoUpperCase != nil {
			r.message = r.ErrorMessageNoUpperCase(inputField, inputValue)
		} else {
			r.message = translations.T("validations.strictpassword.uppercase", helpers.MS{
				":field": inputField,
				":value": inputValue,
			})
		}
		return false
	}

	if r.WithLowerCase && !hasLower {
		if r.ErrorMessageNoLowerCase != nil {
			r.message = r.ErrorMessageNoLowerCase(inputField, inputValue)
		} else {
			r.message = translations.T("validations.strictpassword.lowercase", helpers.MS{
				":field": inputField,
				":value": inputValue,
			})
		}
		return false
	}

	if r.WithDigit && !hasDigit {
		if r.ErrorMessageNoDigit != nil {
			r.message = r.ErrorMessageNoDigit(inputField, inputValue)
		} else {
			r.message = translations.T("validations.strictpassword.digit", helpers.MS{
				":field": inputField,
				":value": inputValue,
			})
		}
		return false
	}

	return true
}
