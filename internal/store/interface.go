package store

type StoreInterface interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Delete(string) error
	Pop(string) (interface{}, error)
	AllKeys() []string
	GetData() map[string]interface{}
}
