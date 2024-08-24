package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		switch input {
		case "exit 0\n":
			{
				os.Exit(0)
			}
		default:
			{
				fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
			}
		}
	}
}
