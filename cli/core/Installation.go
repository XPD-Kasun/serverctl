package core

type Provider interface {
	Exists() (bool, error)
	Install(args []string) error
	Remove(fully bool) error
	Run(command string) error
	Start() error
	Stop() error
	Restart() error
	//Config
}

type Engine struct {
	ConfigPath     string
	OS             OS
	PackageManager *PackageManager
}

func (engine *Engine) Init() {
	engine.OS = DetectOS()
	engine.PackageManager = DetectPkgManager()
}
