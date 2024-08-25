package commands

import (
	"fmt"
	"strings"
)

type Echo struct{}

func NewEcho() *Echo {
	return &Echo{}
}

func (c *Echo) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func (c *Echo) Help() {
	fmt.Println("echo: display a line of text")
}
