package helpers

import (
	"os"
	"reflect"
	"testing"

	"github.com/vector-ops/mapil/database"
)

func TestSerializeData(t *testing.T) {
	data := []database.KeyValue{
		database.ListType{
			Key: "list1",
			Value: []string{
				"value1",
				"value2",
			},
		},
		database.ValueType{
			Key:   "key1",
			Value: "value1",
		},
	}

	expected := `[{"type":"list","data":{"key":"list1","value":["value1","value2"]}},{"type":"value","data":{"key":"key1","value":"value1"}}]`

	jBuf, err := Serialize(data)
	if err != nil {
		t.Fatal(err)
	}
	if jBuf == nil {
		t.Fatalf("Nothing in buffer.")
	}
	if string(jBuf) != expected {
		t.Fatalf("Data does not match.\nexpected: %s\ngot: %s", expected, string(jBuf))
	}
}

func TestDeserializeData(t *testing.T) {
	jBuf := []byte(`[{"type":"list","data":{"key":"list1","value":["value1","value2"]}},{"type":"value","data":{"key":"key1","value":"value1"}}]`)

	expected := []database.KeyValue{
		database.ListType{
			Key:   "list1",
			Value: []string{"value1", "value2"},
		},
		database.ValueType{
			Key:   "key1",
			Value: "value1",
		},
	}

	kv, err := Deserialize(jBuf)
	if err != nil {
		t.Fatal(err)
	}
	if len(kv) == 0 {
		t.Fatal("failed to unmarshall")
	}
	if !reflect.DeepEqual(expected, kv) {
		t.Fatalf("structs not equal\ngot: %+v\nexpected: %+v", kv, expected)
	}
}

func TestDeserializeFile(t *testing.T) {
	fpath := "../mapil/mapil.json"

	file, err := os.Open(fpath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

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

	got, err := DeserializeFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %+v\nGot: %+v", expected, got)
	}
}
