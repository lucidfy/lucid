package lang

import (
	"os"
	"strings"

	"github.com/lucidfy/lucid/pkg/helpers"
)

type Translations struct {
	langs map[string]helpers.MS
}

func Load(langs map[string]helpers.MS) *Translations {
	return &Translations{langs: langs}
}

// T translates based on the default language inside .env
// or if not givem or such not loaded
func (t Translations) Get(key string, values helpers.MS) string {
	lang := os.Getenv("APP_LANGUAGE")
	if lang == "" {
		lang = "en-US"
	}

	return t.GetLang(key, values, lang)
}

func (t Translations) GetLang(key string, values helpers.MS, lang string) string {
	sentence := t.langs[lang][key]
	for k, v := range values {
		sentence = strings.Replace(sentence, k, v, -1)
	}
	return sentence
}
