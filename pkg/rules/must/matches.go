package must

import (
	"fmt"
)

type Matches struct {
	CustomErrorMessage func(string, string, string) string
	TargetField        string
	inputValues        map[string]interface{}
}

func (r *Matches) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.TargetField)
	}

	return fmt.Sprintf("%s did not match with %s", inputField, r.TargetField)
}

func (r *Matches) Valid(inputField string, inputValue string) bool {
	return fmt.Sprint(r.inputValues[r.TargetField]) == inputValue
}

func (r *Matches) Inputs(inputs map[string]interface{}) {
	r.inputValues = inputs
}
