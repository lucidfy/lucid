package must

import "fmt"

type Required struct {
	Value interface{}
}

func (r *Required) ErrorMessage(inputField string, inputValue string) string {
	return fmt.Sprintf("%s is required!", inputField)
}

func (r *Required) Valid(inputField string, inputValue string) bool {
	return len(inputValue) > 0
}
