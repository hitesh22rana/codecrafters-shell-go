package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/commands"
)

func main() {
	cmds := commands.NewCommands()

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)

		data := strings.Split(input, " ")
		cmd := data[0]
		args := data[1:]

		switch cmd {
		case "exit":
			cmds["exit"].Execute(args)

		case "echo":
			cmds["echo"].Execute(args)

		case "type":
			cmds["type"].Execute(args)

		default:
			command := exec.Command(cmd, args...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout

			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", cmd)
			}
		}
	}
}
