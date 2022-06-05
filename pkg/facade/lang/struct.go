package lang

import (
	"os"
	"strings"
)

type Translations struct {
	language       string
	languages_data map[string]map[string]string
}

func Load(languages_data map[string]map[string]string) *Translations {
	return &Translations{languages_data: languages_data}
}

func (t *Translations) SetLanguage(language string) *Translations {
	t.language = language
	return t
}

// T translates based on the default language inside .env
// or if not givem or such not loaded
func (t Translations) Get(key string, values map[string]string) string {
	lang := "en-US"

	if t.language != "" {
		lang = t.language
	} else if os.Getenv("APP_LANGUAGE") != "" {
		lang = os.Getenv("APP_LANGUAGE")
	}

	return t.Direct(lang, key, values)
}

func (t Translations) Direct(lang string, key string, values map[string]string) string {
	sentence := t.languages_data[lang][key]

	for k, v := range values {
		sentence = strings.Replace(sentence, k, v, -1)
	}

	// return the key if the sentence is empty
	if sentence == "" {
		return key
	}

	return sentence
}
