package cookie

type CookieContract interface {
	CreateSessionCookie() interface{}
	Set(string, string) (bool, error)
	Get(string) (interface{}, error)
	Expire(string)
}
