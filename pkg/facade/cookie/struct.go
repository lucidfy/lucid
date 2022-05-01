package cookie

type CookieContract interface {
	Set(name string, value string) (bool, error)
	Get(name string) (*string, error)
}
