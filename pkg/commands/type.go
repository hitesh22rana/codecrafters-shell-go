package commands

import (
	"fmt"
	"os"
	"strings"
)

var cmdsType = make(map[string]string)

func init() {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if _, ok := cmdsType[file.Name()]; !ok {
				cmdsType[file.Name()] = path + "/" + file.Name()
			}
		}
	}
}
func Type(args []string) error {
	for _, command := range args {
		if command == "echo" || command == "exit" || command == "type" {
			fmt.Printf("%s is a shell builtin\n", command)
		} else if path, ok := cmdsType[command]; ok {
			fmt.Printf("%s is %s\n", command, path)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
	return nil
}
