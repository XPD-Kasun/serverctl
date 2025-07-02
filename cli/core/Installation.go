package core

type ExecutableType string

const (
	Standalone     ExecutableType = "Standalone"
	Script         ExecutableType = "Script"
	UnknownExeType ExecutableType = ""
)

type Installer struct {
	os       OS
	file     string
	fileType ExecutableType
}

func NewEmptyInstaller() *Installer {

	return &Installer{
		os:       DetectOS(),
		file:     "",
		fileType: UnknownExeType,
	}
}

func NewInstaller(executable string, exeType ExecutableType) *Installer {

	return &Installer{
		os:       DetectOS(),
		file:     executable,
		fileType: exeType,
	}
}

func (i *Installer) SetExecutable(executable string, executableType ExecutableType) {
	i.file = executable
	i.fileType = executableType
}

func (i *Installer) GetExecutable() string {
	return i.file
}

func (i *Installer) GetExecutableType() ExecutableType {
	return i.fileType
}
