package engines

type EngineContract interface {
	GetRequest() interface{}
	GetResponse() interface{}
	GetUrl() interface{}
}
