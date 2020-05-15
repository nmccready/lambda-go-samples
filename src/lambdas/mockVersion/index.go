package main

import (
	"fmt"

	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

func main() {
	fmt.Println("From -ldflags utils.Version: ", utils.Version)

	utils.GetVersionMut(&utils.Version)
	// make sure utils.DEFAULT_VERSION is not mutated
	fmt.Println("utils.DEFAULT_VERSION: ", utils.DEFAULT_VERSION)

	fmt.Println("Resolved Local or pre utils.Version ", utils.Version)
}
