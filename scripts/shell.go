package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-la")
	out, err := cmd.Output()

	// cmd.Stdin = os.Stdin
	// cmd.Stderr = os.Stderr

	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	if err != nil {
		fmt.Printf("Error %s", err.Error())
		return
	}
	fmt.Println(string(out))

}
