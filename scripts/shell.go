package main

import (
	"fmt"
	"os"

	utils "github.com/nmccready/lambda-go-samples/src/utils"
)

func main() {
	// fmt.Println(os.Args)
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// fmt.Println("path: ", path)
	out, err := utils.Ls(path)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
	fmt.Println(string(out))
}
