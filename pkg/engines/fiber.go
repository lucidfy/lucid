// TODO: replace this with proper fiber implementation!
package engines

type FiberEngine struct {
	HttpResponseWriter interface{}
	HttpRequest        interface{}
	Response           interface{}
	Request            interface{}
	Session            interface{}
	Url                interface{}
}

// ---

func Fiber(w interface{}, r interface{}) FiberEngine {
	return FiberEngine{
		HttpResponseWriter: nil,
		HttpRequest:        nil,
		Response:           nil,
		Request:            nil,
		Session:            nil,
		Url:                nil,
	}
}

func (f FiberEngine) GetRequest() interface{} {
	return f.Request
}
func (f FiberEngine) GetResponse() interface{} {
	return f.Response
}
func (f FiberEngine) GetSession() interface{} {
	return f.Session
}
func (f FiberEngine) GetUrl() interface{} {
	return f.Url
}
