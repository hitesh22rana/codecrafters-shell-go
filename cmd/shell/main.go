package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/commands"
	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
)

func main() {
	trace := utils.NewTrace()
	track := utils.NewTrack()
	cmds := commands.NewCommands(trace, track)

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

		case "help":
			if len(args) != 1 {
				fmt.Println("help: expected 1 argument; got", len(args))
				continue
			}

			if cmd, ok := cmds[args[0]]; ok {
				cmd.Help()
			} else if args[0] == "help" {
				fmt.Println("help: display help for a command")
			} else {
				fmt.Printf("%s: command not found\n", args[0])
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
