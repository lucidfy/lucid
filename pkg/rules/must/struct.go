package must

import (
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type SetOfRules map[string][]Rule

type Rule interface {
	ErrorMessage(inputField string, inputValue string) string
	Valid(inputField string, inputValue string) bool
	SetTranslation(trans *lang.Translations)
}

var TestLanguages = map[string]helpers.MS{
	"en-US": map[string]string{
		"validations.email":                    ":field is not a valid email address!",
		"validations.matches":                  ":field did not match with :target!",
		"validations.max":                      ":field is set to maximum of :length length!",
		"validations.min":                      ":field is set to minimum of :length length!",
		"validations.required":                 ":field is required!",
		"validations.strictpassword.special":   ":field should contain at least 1 special character!",
		"validations.strictpassword.uppercase": ":field should contain at least 1 upper case character!",
		"validations.strictpassword.lowercase": ":field should contain at least 1 lower case character!",
		"validations.strictpassword.digit":     ":field should contain at least 1 digit!",
	},
}
