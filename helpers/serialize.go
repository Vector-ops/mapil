package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/vector-ops/mapil-cli/types"
)

func Serialize(do []types.DataObject, file *os.File) error {
	file.Truncate(int64(file.Fd()))
	enc := json.NewEncoder(file)
	err := enc.Encode(do)
	if err != nil {
		return err
	}
	return nil
}

func Deserialize(file *os.File, do *[]types.DataObject) error {
	dec := json.NewDecoder(file)
	err := dec.Decode(do)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
