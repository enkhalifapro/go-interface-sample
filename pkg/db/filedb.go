package db

import "fmt"

type FileDBProvider struct {
	fileReader FileReader
}

func NewFileDBProvider(fileReader FileReader) *FileDBProvider {
	return &FileDBProvider{fileReader: fileReader}
}

type FileReader interface {
	Read(name string) ([]byte, error)
}

func (f *FileDBProvider) List(collectionName string) (map[string]interface{},error) {
	//os.ReadFile("./vendors.json")
	f.fileReader.Read(fmt.Sprintf("%s.json", collectionName))
	return nil, nil
}
