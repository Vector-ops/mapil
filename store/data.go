package store

import "github.com/vector-ops/mapil-cli/types"

type Data struct {
	List map[string]any `json:"list,omitempty"`
}

func NewData() *Data {
	return &Data{
		List: make(map[string]any),
	}
}

func (d *Data) AddObject(key string, value any) {
	d.List[key] = value
}

func (d *Data) UpdateObject(key string, value any) {
	_, ok := d.List[key]
	if ok {
		d.List[key] = value
	}
}

func (d *Data) GetObject(key string) types.DataObject {
	v, ok := d.List[key]
	if ok {
		return types.DataObject{Key: key, Value: v}
	}
	return types.DataObject{}
}

func (d *Data) GetValue(key string) any {
	v, ok := d.List[key]
	if ok {
		return v
	}
	return nil
}

func (d *Data) GetAllObjects() []types.DataObject {
	var objs []types.DataObject
	for k, v := range d.List {
		objs = append(objs, types.DataObject{Key: k, Value: v})
	}
	return objs
}

func (d *Data) GetAllKeys() []string {
	var keys []string
	for k := range d.List {
		keys = append(keys, k)
	}
	return keys
}

func (d *Data) DeleteObject(key string) {
	delete(d.List, key)
}
