package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
)

const DEFAULT_VERSION = "development"

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided          = errors.New("no name was provided in the HTTP body")
	ErrInvalidGetRequest        = errors.New("invalid GET request")
	Version              string = DEFAULT_VERSION
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

func FullTernary(e bool, a, b interface{}) interface{} {
	if e {
		if reflect.TypeOf(a).Kind() == reflect.Func {
			return a.(func() interface{})()
		}
		return a
	}
	if reflect.TypeOf(b).Kind() == reflect.Func {
		return b.(func() interface{})()
	}
	return b
}

func Ls(toLs string) (string, error) {
	dir := "/tmp"
	if toLs != "" {
		dir = toLs
	}

	cmd := exec.Command("ls", "-la", dir)
	outBytes, err := cmd.Output()
	return string(outBytes), err
}
