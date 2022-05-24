package rules

import (
	"fmt"
	"sync"

	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type Validator struct {
	Translation *lang.Translations
	InputValues map[string]interface{}
}

func New(t *lang.Translations, values map[string]interface{}) *Validator {
	return &Validator{
		Translation: t,
		InputValues: values,
	}
}

func (v *Validator) Validate(
	inputField string,
	rule must.Rule,
	err chan map[string]string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	inputValue := fmt.Sprint(v.InputValues[inputField])

	i, ok := rule.(interface{ Inputs(map[string]interface{}) })
	if ok {
		i.Inputs(v.InputValues)
	}

	rule.SetTranslation(v.Translation)

	if !rule.Valid(inputField, inputValue) {
		err <- map[string]string{
			inputField: rule.ErrorMessage(inputField, inputValue),
		}
	}
}

func (v *Validator) GetErrors(setOfRules *must.SetOfRules) map[string]interface{} {
	var errsChan = make(chan map[string]string)

	var wg sync.WaitGroup

	for inputField, inputRules := range *setOfRules {
		for _, inputRule := range inputRules {
			wg.Add(1)
			go v.Validate(
				inputField,
				inputRule,
				errsChan,
				&wg,
			)
		}
	}

	go func() {
		wg.Wait()
		close(errsChan)
	}()

	validationErrors := map[string]interface{}{}
	for val := range errsChan {
		for k, v := range val {
			validationErrors[k] = v
		}
	}

	return validationErrors
}
