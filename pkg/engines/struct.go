package engines

type EngineInterface interface {
	ParsedRequest() interface{}
	ParsedResponse() interface{}
	ParsedSession() interface{}
}
