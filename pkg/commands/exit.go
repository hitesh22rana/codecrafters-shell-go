package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type Exit struct {
	trace trace.Trace
}

func NewExit(trace trace.Trace) *Exit {
	return &Exit{trace: trace}
}

func (c *Exit) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("exit: too many arguments")
	}

	code, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("exit: %s", err)
	}

	os.Exit(code)
	return nil
}
