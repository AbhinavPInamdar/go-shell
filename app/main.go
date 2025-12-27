package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var input string
	for true {
		fmt.Print("$ ")
		fmt.Scan(&input)
		if input != "cd" || input != "echo" {
			fmt.Println(input, " : command not found")
		}
	}

}
