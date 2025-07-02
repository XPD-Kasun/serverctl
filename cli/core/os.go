package core

import "runtime"

type OSTYPE string
type DistroType string

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
