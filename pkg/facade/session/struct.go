package session

type SesionContract interface {
	Set(name string, value string) (bool, error)
	Get(name string) (*string, error)
	SetFlash(name string, value string)
	GetFlash(name string) *string
	SetFlashMap(name string, values interface{})
	GetFlashMap(name string) *map[string]interface{}
}
