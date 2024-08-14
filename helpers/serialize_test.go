package helpers

import (
	"fmt"
	"testing"

	"github.com/vector-ops/mapil/database"
)

func TestSerializeData(t *testing.T) {
	data := []database.KeyValue{
		database.ListType{
			Key: "lsit1",
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

	expected := `{"key":"list1","value":["value1","value2"]}{"key":"key1","value":"value1"}`

	jBuf, err := Serialize(data)
	if err != nil {
		t.Fatal(err)
	}
	if jBuf == nil {
		t.Fatalf("Nothing in buffer.")
	}
	fmt.Println(string(jBuf))
	fmt.Println(expected)
	if string(jBuf) != expected {
		t.Fatalf("Data does not match.")
	}
}

func TestDeserializeData(t *testing.T) {
	data := []database.KeyValue{
		database.ListType{
			Key: "lsit1",
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

	jBuf, err := Serialize(data)
	if err != nil {
		t.Fatal(err)
	}

	kv, err := Deserialize(jBuf)
	if err != nil {
		t.Fatal(err)
	}
	if len(kv) == 0 {
		t.Fatal("failed to unmarshall")
	}
}
