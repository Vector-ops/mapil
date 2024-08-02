package store

import (
	"fmt"

	"github.com/vector-ops/mapil/helpers"
	"github.com/vector-ops/mapil/types"
)

type Store struct {
	data *Data
	file *helpers.File
}

func NewStore() *Store {
	return &Store{
		data: NewData(),
		file: helpers.NewFileObject(),
	}
}

func (s *Store) Init() {
	s.file.Init()
	s.LoadData()
}

func (s *Store) AddValue(key string, value any) {
	if v := s.data.GetValue(key); v == nil {
		s.data.AddObject(key, value)
	}
}

func (s *Store) UpdateValue(key string, value any) {
	s.data.UpdateObject(key, value)
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

func (s *Store) GetAllData() []types.DataObject {
	return s.data.GetAllObjects()
}

func (s *Store) LoadData() {
	data, err := s.file.LoadFile()
	if err != nil {
		fmt.Println("Failed to load data file.")
	}
	for _, v := range data {
		s.data.AddObject(v.Key, v.Value)
	}
}

func (s *Store) Persist() {
	err := s.file.SaveFile(s.GetAllData())
	if err != nil {
		fmt.Println("Failed to save file")
	}
}
