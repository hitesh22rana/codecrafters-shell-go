package commands

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type Echo struct {
	trace trace.Trace
}

func NewEcho(trace trace.Trace) *Echo {
	return &Echo{trace: trace}
}

func (c *Echo) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
