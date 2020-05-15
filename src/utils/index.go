package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func Dirname__() string {
	_, thisFileName, _, _ := runtime.Caller(1)
	return path.Dir(thisFileName)
}

func GetFileBytes(filename string) []byte {
	filepath := path.Join(Dirname__(), filename)
	// fmt.Println("filepath: " + filepath)
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	return bytes
}
