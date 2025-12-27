package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		command := strings.Split(input, " ")

		switch command[0] {
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		case "exit":
			return
		case "type":
			if command[1] == "echo" || command[1] == "exit" || command[1] == "type" {
				fmt.Println(command[1] + " is a shell builtin")
			} else {
				fmt.Println(command[1] + ": not found")
			}
		default:
			fmt.Println(command[0] + ": command not found")
		}
	}

}
