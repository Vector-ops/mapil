package helpers

import (
	"encoding/json"

	"github.com/vector-ops/mapil/database"
)

func Serialize(data []database.KeyValue) ([]byte, error) {
	buf := make([]byte, 0)

	for _, kv := range data {
		switch kv.(type) {
		case database.ValueType:
			vbuf, err := json.Marshal(kv)
			if err != nil {
				return nil, err
			}
			buf = append(buf, vbuf...)
		case database.ListType:
			lBuf, err := json.Marshal(kv)
			if err != nil {
				return nil, err
			}
			buf = append(buf, lBuf...)
		}
	}

	return buf, nil
}

func Deserialize(data []byte) ([]database.KeyValue, error) {
	kvData := make([]database.KeyValue, 0)
	err := json.Unmarshal(data, &kvData)
	if err != nil {
		return nil, err
	}

	return kvData, nil
}

// serialization and deseria;ization issue
