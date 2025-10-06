package vsftp

import (
	"cli/core"
	"cli/core/consts"
	"errors"
	"os/exec"
)

func installVsftp(engine *core.Engine) error {

	if engine.PackageManager.Name == consts.APT {
		cmd := exec.Command("apt", "install vsftpd")
		err := cmd.Run()
		if err != nil {
			println("APT Installing failed")

		}
	}

	return errors.New("no package managers found")

}
