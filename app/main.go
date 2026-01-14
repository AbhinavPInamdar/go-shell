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
			fmt.Println(command[0] + ": command not found")
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
		fmt.Errorf("empty command")
	}

	cmd1 := cmd[0]
	ar := cmd[1:]

	path, err := exec.LookPath(cmd1)
	if err != nil {
		return err
	}

	c := exec.Command(path, ar...)
	out, err := c.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)

	return nil
}