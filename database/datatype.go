package database

type KeyValue interface {
	GetKey() string
	GetValue() interface{}
}

type ValueType struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ListType struct {
	Key   string   `json:"key,omitempty"`
	Value []string `json:"value,omitempty"`
}

func (vt ValueType) GetKey() string {
	return vt.Key
}
func (vt ValueType) GetValue() interface{} {
	return vt.Value
}
func (lt ListType) GetKey() string {
	return lt.Key
}
func (lt ListType) GetValue() interface{} {
	return lt.Value
}
