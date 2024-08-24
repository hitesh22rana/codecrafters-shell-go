package commands

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type Cd struct {
	trace trace.Trace
}

func NewCd(trace trace.Trace) *Cd {
	return &Cd{trace: trace}
}

func (c *Cd) Execute(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("cd: expected 1 argument; got %d", len(args))
	}

	dir := args[0]

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("cd: %s: No such file or directory", dir)
	}

	c.trace.SetCurrentDir(dir)

	return nil
}
