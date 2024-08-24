package commands

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/pkg/trace"
)

type Command interface {
	Execute(args []string) error
}

func NewCommands(trace trace.Trace) map[string]Command {
	currentDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	trace.SetCurrentDir(currentDir)

	return map[string]Command{
		"exit":    NewExit(trace),
		"echo":    NewEcho(trace),
		"type":    NewType(trace),
		"pwd":     NewPwd(trace),
		"cd":      NewCd(trace),
		"history": NewHistory(trace),
	}
}
