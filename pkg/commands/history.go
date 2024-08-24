package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type History struct {
	trace trace.Trace
}

func NewHistory(trace trace.Trace) *History {
	return &History{trace: trace}
}

func (c *History) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("history: expected 0 arguments; got %d", len(args))
	}

	history := c.trace.GetHistory()
	for i := len(history) - 2; i >= 0; i-- {
		fmt.Println(history[i])
	}

	return nil
}
