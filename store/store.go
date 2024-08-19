package store

import (
	"fmt"
	"strings"

	"github.com/vector-ops/mapil/database"
	"github.com/vector-ops/mapil/helpers"
)

type Store struct {
	data *database.Database
	file *helpers.File
}

func NewStore() *Store {
	return &Store{
		data: database.NewDatabase(),
		file: helpers.NewFileObject(),
	}
}

func (s *Store) Init() error {
	s.file.Init()
	return s.LoadData()
}

func (s *Store) AddValue(key string, value string) {
	s.data.AddObject(database.ValueType{Key: key, Value: value})
}

func (s *Store) AddList(key string, value []string) {
	s.data.AddObject(database.ListType{Key: key, Value: value})
}

func (s *Store) UpdateValue(key string, value string) {
	s.data.UpdateObject(database.ValueType{Key: key, Value: value})
}

func (s *Store) UpdateList(key string, value []string) {
	s.data.UpdateObject(database.ListType{Key: key, Value: value})
}

func (s *Store) DeleteValue(key string) {
	s.data.DeleteObject(key)
}

func (s *Store) DeleteAll() {
	keys := s.data.GetAllKeys()
	for _, k := range keys {
		s.data.DeleteObject(k)
	}
}

func (s *Store) GetKeys() []string {
	return s.data.GetAllKeys()
}

type DataObject struct {
	Key   string
	Value string
}

func (s *Store) GetAllData() []DataObject {
	data := s.data.GetAllObjects()
	var do []DataObject
	for _, kv := range data {
		switch kv.(type) {
		case database.ValueType:
			do = append(do, DataObject{
				Key:   kv.GetKey(),
				Value: kv.GetValue().(string),
			})
		case database.ListType:
			do = append(do, DataObject{
				Key:   kv.GetKey(),
				Value: strings.Join(kv.GetValue().([]string), ", "),
			})
		}
	}
	return do
}

func (s *Store) LoadData() error {
	data, err := s.file.LoadFile()
	if err != nil {
		return fmt.Errorf("failed to load data file")
	}
	for _, v := range data {
		s.data.AddObject(v)
	}
	return nil
}

func (s *Store) Persist() error {
	err := s.file.SaveFile(s.data.GetAllObjects())
	if err != nil {
		return fmt.Errorf("failed to save file")
	}
	return nil
}
