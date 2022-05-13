package must

import (
	"fmt"
	"regexp"
)

type Email struct {
	CustomErrorMessage func(string, string) string
}

func (r *Email) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}

	return fmt.Sprintf("%s is not a valid email address!", inputValue)
}

func (r *Email) Valid(inputField string, inputValue string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(inputValue)
}
