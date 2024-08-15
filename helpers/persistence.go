package helpers

import (
	"fmt"

	"os"
	"path"

	"github.com/vector-ops/mapil/database"
)

const (
	dir      = "mapil"
	fileName = "mapil.json"
)

type File struct {
	filePath string
}

func NewFileObject() *File {
	return &File{}
}

func NewFileObjectWithFile(filePath string) *File {
	return &File{
		filePath: filePath,
	}
}

func (f *File) Init() {
	home, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("failed to create data file. ", err)
	}

	dirPath := path.Join(home, dir)

	if f.filePath == "" {
		f.filePath = path.Join(dirPath, fileName)
	}

	if !checkDirExists(dirPath) {
		if err := createDir(dirPath); err != nil {
			fmt.Println("failed to create data file. ", err)
		}
	}
	if err := f.CreateFile(); err != nil {
		fmt.Println("failed to create data file. ", err)
	}
}

func (f *File) CreateFile() error {

	file, err := os.OpenFile(f.filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil

}

func (f *File) SaveFile(data []database.KeyValue) error {
	if err := writeToFile(data, f.filePath); err != nil {
		return err
	}
	return nil
}

func (f *File) LoadFile() ([]database.KeyValue, error) {
	var data []database.KeyValue
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err = DeserializeFile(file)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

/**
Utility Functions
*/

func writeToFile(data []database.KeyValue, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Truncate(int64(file.Fd()))
	b, err := Serialize(data)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
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
