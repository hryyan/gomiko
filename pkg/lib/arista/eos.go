package arista

import (
	"github.com/hryyan/gomiko/pkg/driver"
	"github.com/hryyan/gomiko/pkg/types"
)

type EOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *EOSDevice) Connect() error {
	return d.base.Connect()

}

func (d *EOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *EOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}

func (d *EOSDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *EOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}
