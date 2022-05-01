package session

type SesionContract interface {
	Put(name string, value string) (bool, error)
	Set(name string, value string) (bool, error)
	Get(name string) (*string, error)
	Destroy(name string) error

	SetFlash(name string, value string)
	GetFlash(name string) *string
	SetFlashMap(name string, values interface{})
	GetFlashMap(name string) *map[string]interface{}
}
