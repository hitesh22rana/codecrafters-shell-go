package utils

type Trace interface {
	AddHistory(command string)
	GetHistory() []string
}

type ShellTrace struct {
	history []string
}

func NewTrace() *ShellTrace {
	return &ShellTrace{
		history: []string{},
	}
}

func (m *ShellTrace) AddHistory(command string) {
	m.history = append(m.history, command)
}

func (m *ShellTrace) GetHistory() []string {
	return m.history
}
