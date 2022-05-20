package translations

import (
	"os"
	"strings"

	"github.com/lucidfy/lucid/pkg/helpers"
)

var Translations = map[string]helpers.MS{
	"en-US": EnglishUnitedStates,
	"zh-CN": ChineseSimplified,
	"zh-TW": ChineseTraditional,
}

// T translates based on the default language inside .env
// or if not givem or such not loaded
func T(key string, values helpers.MS) string {
	lang := os.Getenv("APP_LANGUAGE")
	if lang == "" {
		lang = "en-US"
	}

	return TLang(key, values, lang)
}

func TLang(key string, values helpers.MS, lang string) string {
	sentence := Translations[lang][key]
	for k, v := range values {
		sentence = strings.Replace(sentence, k, v, -1)
	}
	return sentence
}
