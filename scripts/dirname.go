package main

import (
	"fmt"

	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

func main() {
	version, _ := utils.GetVersionJson()
	fmt.Println(version)
}
