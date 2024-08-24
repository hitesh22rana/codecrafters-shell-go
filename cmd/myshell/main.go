package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/commands"
)

var cmds = map[string]func([]string) error{
	"exit": commands.Exit,
	"echo": commands.Echo,
	"type": commands.Type,
}

func main() {
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
			cmds["exit"](args)

		case "echo":
			cmds["echo"](args)

		case "type":
			cmds["type"](args)

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
