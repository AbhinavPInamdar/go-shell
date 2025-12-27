package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var input string
	fmt.Print("$ ")
	fmt.Scanln(&input)
	if input != "cd" || input != "echo" {
		fmt.Print("Invalid command")
	}
}
