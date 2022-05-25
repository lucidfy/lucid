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

func Get(key string, values helpers.MS) string {
	return lang.Load(Languages).Get(key, values)
}
