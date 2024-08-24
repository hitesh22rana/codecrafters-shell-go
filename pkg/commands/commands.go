package commands

type Command interface {
	Execute(args []string) error
}

func NewCommands() map[string]Command {
	return map[string]Command{
		"exit": &Exit{},
		"echo": &Echo{},
		"type": &Type{},
		"pwd":  &Pwd{},
	}
}
