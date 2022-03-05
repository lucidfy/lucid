package must

import (
	"fmt"
)

type Min struct {
	Value interface{}
}

func (r *Min) ErrorMessage(inputField string, inputValue string) string {
	return fmt.Sprintf("%s should be minimum of %d", inputField, r.Value.(int))
}

func (r *Min) Valid(inputField string, inputValue string) bool {
	return len(inputValue) >= r.Value.(int)
}
