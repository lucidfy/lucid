package rules

import (
	"sync"

	"github.com/daison12006013/lucid/pkg/rules/must"
)

func Validate(
	inputField string,
	inputValue string,
	rule must.Rule,
	err chan map[string]string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	if !rule.Valid(inputField, inputValue) {
		err <- map[string]string{
			inputField: rule.ErrorMessage(inputField, inputValue),
		}
	}
}
