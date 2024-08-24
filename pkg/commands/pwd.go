package commands

import (
	"fmt"
	"os"
)

type Pwd struct{}

func NewPwd() *Pwd {
	return &Pwd{}
}

func (c *Pwd) Execute(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}
