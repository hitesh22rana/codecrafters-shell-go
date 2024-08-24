package commands

import (
	"fmt"
	"os"
	"path/filepath"

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

	if dir == "~" {
		dir = os.Getenv("HOME")
	}

	if dir[0] != '/' {
		dir = filepath.Join(c.trace.GetCurrentDir(), dir)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("cd: %s: No such file or directory", dir)
	}

	c.trace.SetCurrentDir(dir)

	return nil
}
