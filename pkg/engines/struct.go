package engines

type EngineContract interface {
	GetRequest() interface{}
	GetResponse() interface{}
	GetSession() interface{}
	GetUrl() interface{}
}
