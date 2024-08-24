package trace

type Trace interface {
	SetCurrentDir(dir string)
	GetCurrentDir() string
	AddHistory(command string)
	GetHistory() []string
}

type ShellTrace struct {
	currentDir string
	history    []string
}

func NewTrace() *ShellTrace {
	return &ShellTrace{
		currentDir: "",
		history:    []string{},
	}
}

func (m *ShellTrace) SetCurrentDir(dir string) {
	m.currentDir = dir
}

func (m *ShellTrace) GetCurrentDir() string {
	return m.currentDir
}

func (m *ShellTrace) AddHistory(command string) {
	m.history = append(m.history, command)
}

func (m *ShellTrace) GetHistory() []string {
	return m.history
}
