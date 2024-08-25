package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
)

type Cd struct {
	track utils.Track
}

func NewCd(track utils.Track) *Cd {
	return &Cd{track: track}
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
		dir = filepath.Join(c.track.GetCurrentDir(), dir)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("cd: %s: No such file or directory", dir)
	}

	c.track.SetCurrentDir(dir)

	return nil
}

func (c *Cd) Help() {
	fmt.Println("cd: change the current directory")
}
