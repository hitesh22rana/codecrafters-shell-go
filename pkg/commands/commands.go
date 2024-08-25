package commands

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
)

type Command interface {
	Execute(args []string) error
	Help()
}

func NewCommands(trace utils.Trace, track utils.Track) map[string]Command {
	currentDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	track.SetCurrentDir(currentDir)

	return map[string]Command{
		"exit":    NewExit(),
		"echo":    NewEcho(),
		"type":    NewType(),
		"pwd":     NewPwd(track),
		"cd":      NewCd(track),
		"history": NewHistory(trace),
	}
}
