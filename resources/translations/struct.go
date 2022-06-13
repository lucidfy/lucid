package translations

import (
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/helpers"
)

var Languages = map[string]helpers.MS{
	"en-US": EnglishUnitedStates,
	"zh-CN": ChineseSimplified,
	"zh-TW": ChineseTraditional,
}

func load() *lang.Translations {
	return lang.Load(Languages)
}

func Get(key string, values helpers.MS) string {
	return load().Get(key, values)
}

func Direct(lang string, key string, values helpers.MS) string {
	return load().Direct(lang, key, values)
}
