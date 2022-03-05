package must

import (
	"fmt"
	"regexp"
)

type Email struct{}

func (r *Email) ErrorMessage(inputField string, inputValue string) string {
	return fmt.Sprintf("%s is not a valid email address!", inputField)
}

func (r *Email) Valid(inputField string, inputValue string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(inputValue)
}
