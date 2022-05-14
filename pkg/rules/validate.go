package rules

import (
	"fmt"
	"sync"

	"github.com/lucidfy/lucid/pkg/rules/must"
)

func Validate(
	inputField string,
	inputValues map[string]interface{},
	rule must.Rule,
	err chan map[string]string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	inputValue := fmt.Sprint(inputValues[inputField])

	i, ok := rule.(interface{ Inputs(map[string]interface{}) })
	if ok {
		i.Inputs(inputValues)
	}

	if !rule.Valid(inputField, inputValue) {
		err <- map[string]string{
			inputField: rule.ErrorMessage(inputField, inputValue),
		}
	}
}

func GetErrors(setOfRules *must.SetOfRules, inputValues map[string]interface{}) map[string]interface{} {
	var errsChan = make(chan map[string]string)

	var wg sync.WaitGroup

	for inputField, inputRules := range *setOfRules {
		for _, inputRule := range inputRules {
			wg.Add(1)
			go Validate(
				inputField,
				inputValues,
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
