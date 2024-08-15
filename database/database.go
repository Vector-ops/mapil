package database

import "fmt"

type Database struct {
	List map[string]KeyValue `json:"list,omitempty"`
}

func NewDatabase() *Database {
	return &Database{
		List: make(map[string]KeyValue),
	}
}

func (d *Database) AddObject(kv KeyValue) error {
	if _, ok := d.List[kv.GetKey()]; ok {
		return fmt.Errorf("key already exists")
	}
	d.List[kv.GetKey()] = kv
	return nil
}

func (d *Database) UpdateObject(kv KeyValue) error {
	if _, ok := d.List[kv.GetKey()]; ok {
		d.List[kv.GetKey()] = kv
		return nil
	}
	return fmt.Errorf("key does not exist")
}

func (d *Database) GetObject(key string) (KeyValue, error) {
	if kv, ok := d.List[key]; ok {
		return kv, nil
	}
	return nil, fmt.Errorf("key does not exist")
}

func (d *Database) GetValue(key string) (interface{}, error) {
	if kv, ok := d.List[key]; ok {
		return kv.GetValue(), nil
	}
	return nil, fmt.Errorf("key does not exist")
}

func (d *Database) GetAllObjects() []KeyValue {
	var objs []KeyValue
	for _, kv := range d.List {
		objs = append(objs, kv)
	}
	return objs
}

func (d *Database) GetAllKeys() []string {
	keys := make([]string, 0, len(d.List))
	for k := range d.List {
		keys = append(keys, k)
	}
	return keys
}

func (d *Database) DeleteObject(key string) {
	delete(d.List, key)
}
