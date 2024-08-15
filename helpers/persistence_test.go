package helpers

import (
	"os"
	"reflect"
	"testing"

	"github.com/vector-ops/mapil/database"
)

func TestLoadFile(t *testing.T) {
	file := NewFileObjectWithFile("/home/vector/dev/mapil/mapil/mapil.json")
	file.Init()
	got, err := file.LoadFile()
	if err != nil {
		t.Fatal(err)
	}

	expected := []database.KeyValue{
		database.ValueType{
			Key:   "wall",
			Value: "gang",
		},
		database.ListType{
			Key:   "dell",
			Value: []string{"geng", "random", "list"},
		},
		database.ValueType{
			Key:   "shell",
			Value: "gang",
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected: %+v\ngot: %+v", expected, got)
	}
}

func TestSaveFile(t *testing.T) {
	file := NewFileObjectWithFile("/home/vector/dev/mapil/mapil/mapil.json")
	file.Init()

	data := []database.KeyValue{
		database.ValueType{
			Key:   "test1",
			Value: "testvalue1",
		},
		database.ListType{
			Key:   "dell",
			Value: []string{"geng", "random", "list"},
		},
		database.ValueType{
			Key:   "shell",
			Value: "gang",
		},
	}
	err := file.SaveFile(data)
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(file.filePath)
	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatalf("failed to write to file")
	}
}
