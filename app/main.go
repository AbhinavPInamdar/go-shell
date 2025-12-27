package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var input string
	for true {
		fmt.Println("$ ")
		fmt.Scanln(&input)
		if input != "cd" || input != "echo" {
			fmt.Print(input, ": command not found")
		}
	}

}
