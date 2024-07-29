package types

type DataObject struct {
	Key   string `json:"key,omitempty"`
	Value any    `json:"value,omitempty"`
}
