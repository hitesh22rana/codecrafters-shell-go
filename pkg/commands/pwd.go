package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type Pwd struct {
	trace trace.Trace
}

func NewPwd(trace trace.Trace) *Pwd {
	return &Pwd{trace: trace}
}

func (c *Pwd) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("pwd: expected 0 arguments; got %d", len(args))
	}

	currentDir := c.trace.GetCurrentDir()
	fmt.Println(currentDir)

	return nil
}
