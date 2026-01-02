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
			checkType(command[1])
		default:
			fmt.Println(command[0] + ": command not found")
		}
	}
}

func checkType(cmd string) {
	if cmd == "echo" || cmd == "exit" {
		fmt.Println(cmd + "is a shell builtin")
	}
	path, err := os.LookPath(cmd)
	if err == nil {
		fmt.Prinln(cmd + "is" + path)
	}

	fmt.Println(path + ": not found")
}