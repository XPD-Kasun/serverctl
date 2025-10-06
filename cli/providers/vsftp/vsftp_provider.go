package vsftp

import "cli/core"

type VSFTPProvider struct {
	engine *core.Engine
}

func NewVSFTPProvider() *VSFTPProvider {
	var engine = core.Engine{}
	engine.Init()
	var vsftp *VSFTPProvider = &VSFTPProvider{engine: &engine}
	return vsftp
}

func (vsftp *VSFTPProvider) Exists() (bool, error) {
	return detectVSFtp(vsftp)
}

func (vsftp *VSFTPProvider) Install(args []string) error {
	return installVsftp(vsftp.engine)
}

func (vsftp *VSFTPProvider) Remove(fully bool) error {
	panic("not implemented") // TODO: Implement
}

func (vsftp *VSFTPProvider) Run(command string) error {
	panic("not implemented") // TODO: Implement
}

func (vsftp *VSFTPProvider) Start() error {
	panic("not implemented") // TODO: Implement
}

func (vsftp *VSFTPProvider) Stop() error {
	panic("not implemented") // TODO: Implement
}

func (vsftp *VSFTPProvider) Restart() error {
	panic("not implemented") // TODO: Implement
}
