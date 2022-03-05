package must

import "fmt"

type Max struct {
	Value interface{}
}

func (r *Max) ErrorMessage(inputField string, inputValue string) string {
	return fmt.Sprintf("%s should be maximum of %s", inputField, inputValue)
}

func (r *Max) Valid(inputField string, inputValue string) bool {
	return len(inputValue) <= r.Value.(int)
}
