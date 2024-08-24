package trace

type Trace interface {
	SetCurrentDir(dir string)
	GetCurrentDir() string
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
