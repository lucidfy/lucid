package must

import (
	"fmt"
	"unicode"
)

type StrictPassword struct {
	Value   string
	message string

	WithSpecialChar bool
	WithUpperCase   bool
	WithLowerCase   bool
	WithDigit       bool
}

func (r *StrictPassword) ErrorMessage(inputField string, inputValue string) string {
	if len(r.message) > 0 {
		return r.message
	}
	return "Strict password did not pass the validation!"
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
		r.message = fmt.Sprintf("%s should contain at least 1 special character!", inputField)
		return false
	}

	if r.WithUpperCase && !hasUpper {
		r.message = fmt.Sprintf("%s should contain at least 1 upper case character!", inputField)
		return false
	}

	if r.WithLowerCase && !hasLower {
		r.message = fmt.Sprintf("%s should contain at least 1 lower case character!", inputField)
		return false
	}

	if r.WithDigit && !hasDigit {
		r.message = fmt.Sprintf("%s should contain at least 1 digit!", inputField)
		return false
	}

	return true
}
