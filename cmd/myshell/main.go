package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			code, err := strconv.Atoi(args[0])
			if err != nil {
				os.Exit(1)
			}

			os.Exit(code)

		case "echo":
			fmt.Println(strings.Join(args, " "))

		case "type":
			command := args[0]
			if command == "echo" || command == "exit" || command == "type" {
				fmt.Printf("%s is a shell builtin\n", command)
			} else {
				fmt.Printf("%s: not found\n", command)
			}

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
