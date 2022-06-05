package engines

import (
	"fmt"
	"net/http"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/cookie"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/urls"
)

type NetHttpEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
	Translation    *lang.Translations

	Response response.NetHttpResponse
	Request  request.NetHttpRequest
	URL      urls.NetHttpURL
	Cookie   cookie.NetHttpCookie

	HttpErrorHandler func(EngineContract, *errors.AppError, interface{})
}

func NetHttp(w http.ResponseWriter, r *http.Request, t *lang.Translations) *NetHttpEngine {
	res := response.NetHttp(w, r)
	url := urls.NetHttp(w, r)
	req := request.NetHttp(w, r, t, url)
	coo := cookie.NetHttp(w, r)

	return &NetHttpEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Translation:    t,
		Response:       *res,
		Request:        *req,
		URL:            *url,
		Cookie:         *coo,
	}
}

func (m NetHttpEngine) GetTranslation() *lang.Translations {
	return m.Translation
}

func (m NetHttpEngine) GetRequest() request.RequestContract {
	return &(m.Request)
}

func (m NetHttpEngine) GetResponse() response.ResponseContract {
	return &(m.Response)
}

func (m NetHttpEngine) GetURL() urls.URLContract {
	return &(m.URL)
}

func (m NetHttpEngine) GetCookie() cookie.CookieContract {
	return &(m.Cookie)
}

func (m NetHttpEngine) DD(data ...interface{}) {
	err_msg := fmt.Sprintf("%+v\n", data...)
	m.HttpErrorHandler(m, &errors.AppError{
		Error:   fmt.Errorf("%s", err_msg),
		Message: m.Translation.Get("Die Dump", nil),
		Code:    http.StatusNotImplemented,
	}, nil)
}
