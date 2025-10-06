package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

type OSTYPE string
type DistroType string
type PackageManager struct {
	Name string
	Path string
}

const (
	Windows   OSTYPE = "Windows"
	Linux     OSTYPE = "Linux"
	Darwin    OSTYPE = "Darwin"
	UnknownOS OSTYPE = "UnknownOS"
)
const (
	Win11         DistroType = "Win11"
	Ubuntu        DistroType = "Ubuntu"
	Debian        DistroType = "Debian"
	UnknownDistro DistroType = "UnknownDistro"
)

type OS struct {
	Type   OSTYPE
	Distro DistroType
}

func detectWindowsDistro() DistroType {
	return "dsf"
}

func detectLinuxDistro() DistroType {
	/*
	 * We check the distro type by following order
	 * 1) First check os-release
	 * 2) Check lsb-version
	 *
	 */
	file, error := os.Open("/etc/os-release")
	if error != nil {
		panic("error at " + error.Error())
	}
	content, error := io.ReadAll(file)
	if error != nil {
		panic("error at reading " + error.Error())
	}
	fmt.Println(string(content))
	fmt.Println(exec.LookPath("pacman"))
	return "sdf"
}

func DetectOS() OS {
	currentOS := runtime.GOOS
	switch currentOS {
	case "windows":
		return OS{
			Type:   Windows,
			Distro: detectWindowsDistro(),
		}
	case "linux":
		return OS{
			Type:   Linux,
			Distro: detectLinuxDistro(),
		}
	}
	return OS{
		Type:   UnknownOS,
		Distro: UnknownDistro,
	}

}

func DetectPkgManager() *PackageManager {

	var pacman PackageManager
	var knownPackageManagers = []string{"apt", "rpm", "pacman", "apk", "dnf", "powershell", "cmd"}
	found := false

	for _, tool := range knownPackageManagers {

		path, err := exec.LookPath(tool)
		if err != nil {
			continue
		}
		pacman.Name = tool
		pacman.Path = path
		found = true
		break
	}

	if found {
		return &pacman
	}
	return nil
}
