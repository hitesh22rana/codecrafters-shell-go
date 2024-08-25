package commands

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/pkg/utils"
)

type Pwd struct {
	track utils.Track
}

func NewPwd(track utils.Track) *Pwd {
	return &Pwd{track: track}
}

func (c *Pwd) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("pwd: expected 0 arguments; got %d", len(args))
	}

	currentDir := c.track.GetCurrentDir()
	fmt.Println(currentDir)

	return nil
}

func (c *Pwd) Help() {
	fmt.Println("pwd: output the current working directory")
}
