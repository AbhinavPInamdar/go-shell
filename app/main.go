package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var input string
	i := true
	for i {
		fmt.Print("$ ")
		fmt.Scan(&input)
		if input != "cd" && input != "echo" && input != "break" {
			fmt.Println(input + ": command not found")
		} else if input == "break" {
			i = false
			break
		}
	}

}
