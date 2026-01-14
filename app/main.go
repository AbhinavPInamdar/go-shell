package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			runProgram(command)
		}
	}
}

func checkType(cmd string) {
	if cmd == "echo" || cmd == "exit" || cmd == "type" {
		fmt.Println(cmd + " is a shell builtin")
		return 
	}
	path, err := exec.LookPath(cmd)
	if err == nil {
		fmt.Println(cmd + " is " + path)
		return 
	}

	fmt.Println(cmd + ": not found")
}

func runProgram(cmd []string) error {
	if len(cmd) == 0 {
		return nil
	}

	path, err := exec.LookPath(cmd[0])
	if err != nil {
		fmt.Println(cmd[0] + ": not found")
		return err
	}

	c := exec.Command(path, cmd[1:]...)
	c.Args = cmd
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	
	return c.Run()
}