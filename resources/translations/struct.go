package translations

import (
	"os"
	"strings"

	"github.com/lucidfy/lucid/pkg/helpers"
)

var Translations = map[string]helpers.MS{
	"en-US": EnglishUnitedStates,
}

// T translates based on the default language
func T(key string, values helpers.MS) string {
	return TLang(key, values, os.Getenv("APP_LANGUAGE"))
}

func TLang(key string, values helpers.MS, lang string) string {
	sentence := Translations[lang][key]
	for k, v := range values {
		sentence = strings.Replace(sentence, k, v, -1)
	}
	return sentence
}
