package cisco

import (
	"github.com/hryyan/gomiko/pkg/driver"
	"github.com/hryyan/gomiko/pkg/types"
)

type NXOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *NXOSDevice) Connect() error {
	return d.base.Connect()

}

func (d *NXOSDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *NXOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *NXOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
func (d *NXOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}
