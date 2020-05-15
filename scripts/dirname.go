package main

import (
	"fmt"

	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

var Version string

func main() {
	utils.GetVersionMut(&Version)
	fmt.Println(Version)
}
