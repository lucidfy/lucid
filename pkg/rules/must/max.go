package must

import (
	"fmt"
	"strconv"
)

type Max struct {
	CustomErrorMessage func(string, string) string
	Value              int
}

func (r *Max) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}

	return fmt.Sprintf("%s is set to maximum of %s length", inputField, strconv.Itoa(r.Value))
}

func (r *Max) Valid(inputField string, inputValue string) bool {
	return len(inputValue) <= r.Value
}
