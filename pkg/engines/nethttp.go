package engines

import (
	e "errors"
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/cookie"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/request"
	"github.com/lucidfy/lucid/pkg/facade/response"
	"github.com/lucidfy/lucid/pkg/facade/session"
	"github.com/lucidfy/lucid/pkg/facade/urls"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type NetHttpEngine struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
	Translation    *lang.Translations

	Response response.NetHttpResponse
	Request  request.NetHttpRequest
	URL      urls.NetHttpURL
	Cookie   cookie.NetHttpCookie
	Session  session.SessionContract
}

func NetHttp(w http.ResponseWriter, r *http.Request, t *lang.Translations) *NetHttpEngine {
	res := response.NetHttp(w, r)
	url := urls.NetHttp(w, r)
	req := request.NetHttp(w, r, t, url)
	coo := cookie.NetHttp(w, r)
	var ses session.SessionContract

	if helpers.IsTrue(helpers.Getenv("SESSION_ENABLED", "true")) {
		sessionKey, app_err := coo.Get(helpers.Getenv("SESSION_NAME", "lucid_session"))
		if app_err != nil && e.Is(app_err.Error, http.ErrNoCookie) {
			sessionKey = coo.CreateSessionCookie()
		}

		ses = session.Driver(
			helpers.Getenv("SESSION_DRIVER", "file"),
			sessionKey.(string),
		)
	}

	return &NetHttpEngine{
		ResponseWriter: w,
		HttpRequest:    r,
		Translation:    t,
		Response:       *res,
		Request:        *req,
		URL:            *url,
		Cookie:         *coo,
		Session:        ses,
	}
}

func (m NetHttpEngine) GetRequest() interface{} {
	return m.Request
}

func (m NetHttpEngine) GetResponse() interface{} {
	return m.Response
}

func (m NetHttpEngine) GetURL() interface{} {
	return m.URL
}
