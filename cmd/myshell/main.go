package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/commands"
	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

func main() {
	trace := trace.NewTrace()
	cmds := commands.NewCommands(trace)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		trace.AddHistory(input)

		data := strings.Split(input, " ")
		cmd := data[0]
		args := data[1:]

		switch cmd {
		case "exit":
			err := cmds["exit"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

		case "echo":
			err := cmds["echo"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

		case "type":
			err := cmds["type"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

		case "pwd":
			err := cmds["pwd"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

		case "cd":
			err := cmds["cd"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

		case "history":
			err := cmds["history"].Execute(args)
			if err != nil {
				fmt.Println(err)
			}

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
