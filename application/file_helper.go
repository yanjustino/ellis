package application

import (
	"io/ioutil"
	"os"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadFile(path string) []byte {
	if FileExists(path) {
		value, err := ioutil.ReadFile(path)
		HandleError(err)
		return value
	}
	return nil
}

func RemoveFile(path string) {
	if FileExists(path) {
		e := os.Remove(path)
		HandleError(e)
	}
}

func OpenFile(path string) *os.File {
	file, e := os.Open(path)
	HandleError(e)
	return file
}

func CreateFile(fileName string, content []byte) error {
	return ioutil.WriteFile(fileName, content, 0644)
}
