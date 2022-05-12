package must

import "fmt"

type Required struct {
	CustomErrorMessage func(string, string) string
	Value              interface{}
}

func (r *Required) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue)
	}
	return fmt.Sprintf("%s is required!", inputField)
}

func (r *Required) Valid(inputField string, inputValue string) bool {
	return len(inputValue) > 0
}
