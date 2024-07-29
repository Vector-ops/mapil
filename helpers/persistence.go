package helpers

import (
	"fmt"

	"os"
	"path"

	"github.com/vector-ops/mapil/types"
)

const (
	dir      = "mapil"
	fileName = "mapil.json"
)

var filePath string

type File struct{}

func NewFileObject() *File {
	return &File{}
}

func (f *File) Init() {
	if err := f.CreateFile(); err != nil {
		fmt.Println("failed to create data file. ", err)
	}
}

func (f *File) CreateFile() error {
	home, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	dirPath := path.Join(home, dir)

	filePath = path.Join(dirPath, fileName)

	if !checkDirExists(dirPath) {
		if err := createDir(dirPath); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil

}

func (f *File) SaveFile(data []types.DataObject) error {
	if err := writeToFile(data); err != nil {
		return err
	}
	return nil
}

func (f *File) LoadFile() ([]types.DataObject, error) {
	var data []types.DataObject
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = Deserialize(file, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

/**
Utility Functions
*/

func writeToFile(data []types.DataObject) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	err = Serialize(data, file)
	if err != nil {
		return err
	}

	return nil
}

func createDir(dirPath string) error {
	err := os.Mkdir(dirPath, 0777)
	if err != nil {
		return err
	}

	return nil
}

func checkDirExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return err == nil
}
