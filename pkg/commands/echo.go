package commands

import (
	"fmt"
	"strings"
)

type Echo struct{}

func (c *Echo) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
