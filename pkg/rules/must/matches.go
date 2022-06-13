package must

import (
	"fmt"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type Matches struct {
	CustomErrorMessage func(string, string, string) string
	TargetField        string
	inputValues        map[string]interface{}
	Translation        *lang.Translations
}

func (r *Matches) ErrorMessage(inputField string, inputValue string) string {
	if r.CustomErrorMessage != nil {
		return r.CustomErrorMessage(inputField, inputValue, r.TargetField)
	}
	return r.Translation.Get("validations.matches", helpers.MS{
		":field":  inputField,
		":value":  inputValue,
		":target": r.TargetField,
	})
}

func (r *Matches) Valid(inputField string, inputValue string) bool {
	return fmt.Sprint(r.inputValues[r.TargetField]) == inputValue
}

func (r *Matches) Inputs(inputs map[string]interface{}) {
	r.inputValues = inputs
}

func (r *Matches) SetTranslation(t *lang.Translations) {
	r.Translation = t
}
