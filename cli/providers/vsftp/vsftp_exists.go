package vsftp

import (
	"cli/core/consts"
	"errors"
	"os/exec"
)

func detectVSFtp(vsftp *VSFTPProvider) (bool, error) {
	if vsftp.engine.PackageManager.Name == consts.APT {
		out, err := exec.Command("dpkg-query", "-W", "-f='${Package}'", "vsftpd").CombinedOutput()
		if err != nil {
			return false, errors.New(string(out))
		}
		return true, nil
	}
	return false, nil
}
