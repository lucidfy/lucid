package must

import (
	"fmt"
	"strconv"
)

type Min struct {
	CustomErrorMessage func(string, string) string
	Value              int
}

func (r *Min) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}

	return fmt.Sprintf("%s is set to minimum of %s length", inputField, strconv.Itoa(r.Value))
}

func (r *Min) Valid(inputField string, inputValue string) bool {
	return len(inputValue) >= r.Value
}
