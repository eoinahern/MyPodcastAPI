package util

import (
	"log"
	"os"
)

type FileHelperUtil struct {
	file []byte
}

func createDIR() {

}

func (f *FileHelperUtil) CheckDirFileExists(path string) bool {

	_, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	}

	return true
}

func (f *FileHelperUtil) CheckExtension() bool {
	return false
}

func (f *FileHelperUtil) CheckMaxSizeAllowed() {

}

func (f *FileHelperUtil) CreateDir(dirpath string) {

	err := os.Mkdir(dirpath, 777)

	if err != nil {
		log.Println(err)
	}
}
