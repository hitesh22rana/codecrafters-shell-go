package utils

type Track interface {
	SetCurrentDir(dir string)
	GetCurrentDir() string
}

type ShellTrack struct {
	currentDir string
}

func NewTrack() *ShellTrack {
	return &ShellTrack{
		currentDir: "",
	}
}

func (m *ShellTrack) SetCurrentDir(dir string) {
	m.currentDir = dir
}

func (m *ShellTrack) GetCurrentDir() string {
	return m.currentDir
}
