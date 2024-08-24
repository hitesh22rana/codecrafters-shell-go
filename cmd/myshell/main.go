package main

import (
	"bufio"
	"fmt"
	"os"
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
			{
				if len(args) > 0 && strings.Join(args, " ") == "0" {
					os.Exit(0)
				}
			}

		case "echo":
			fmt.Println(strings.Join(args, " "))

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
