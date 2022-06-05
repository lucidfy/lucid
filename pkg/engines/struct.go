package engines

import (
	"github.com/lucidfy/lucid/pkg/facade/cookie"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/urls"
)

type EngineContract interface {
	GetRequest() request.RequestContract
	GetResponse() response.ResponseContract
	GetURL() urls.URLContract
	GetCookie() cookie.CookieContract
	GetTranslation() *lang.Translations
	DD(data ...interface{})
}
