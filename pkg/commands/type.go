package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Type struct{}

func NewType() *Type {
	return &Type{}
}

func (c *Type) Execute(args []string) error {
	for _, command := range args {
		switch command {
		case "echo", "exit", "type", "pwd":
			fmt.Printf("%s is a shell builtin\n", command)
		default:
			paths := strings.Split(os.Getenv("PATH"), ":")
			var isFound bool = false
			for _, path := range paths {
				fullPath := filepath.Join(path, command)
				if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
					fmt.Printf("%s is %s\n", command, fullPath)
					isFound = true
					break
				}
			}

			if !isFound {
				fmt.Printf("%s: not found\n", command)
			}
		}
	}

	return nil
}
