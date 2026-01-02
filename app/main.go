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

func checkType(path string) {
	if path == "echo" || path == "exit" {
		fmt.Println(path + "is a shell builtin")
	}
	pathValue := os.Getenv("PATH")
	dirs := strings.Split(pathValue, ":")
	for _, dir := range dirs {
		if dir != "" {
			continue
		}

		fullpath := dir + "/" + path
		fileInfo, err := os.Stat(fullpath)
		if err == nil && !fileInfo.IsDir() {
			fmt.Println(path + "is" + fullpath )
			return 
		}

	}
	fmt.Println(path + ": not found")
}