package store

type StoreInterface interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Delete(string) error
	Pop() (interface{}, bool)
	AllKeys() []string
	GetData() map[string]interface{}
}
