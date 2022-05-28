package engines

type EngineContract interface {
	GetRequest() interface{}
	GetResponse() interface{}
	GetURL() interface{}
	GetCookie() interface{}
	GetSession() interface{}
}
