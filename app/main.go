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
		if command[0] != "cd" && command[0] != "echo" && command[0] != "exit" {
			fmt.Println(command[0] + ": command not found")
		} else if command[0] == "exit" {
			break
		} else if command[0] == "echo" {
			fmt.Println(strings.Join(command[1:], " "))
		}

	}

}
