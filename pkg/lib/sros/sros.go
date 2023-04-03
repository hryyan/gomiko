package sros

import (
	"errors"

	"github.com/hryyan/gomiko/pkg/driver"
)

type SROSDevice struct {
	Driver     driver.IDriver
	DeviceType string
	Prompt     string
}

func (d *SROSDevice) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\\*?([ABCD]:\\S*@?\\S+)[#>%]", "#")
	if err != nil {
		return err
	}
	d.Prompt = prompt

	return d.sessionPreparation()
}

func (d *SROSDevice) Disconnect() {

	d.Driver.Disconnect()

}

func (d *SROSDevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *SROSDevice) SendConfigSet(cmds []string) (string, error) {

	cmds = append(cmds, "logout")
	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	return out, err

}

func (d *SROSDevice) sessionPreparation() error {
	if _, err := d.SendCommand("environment no more"); err != nil {
		return errors.New("failed to disable pagination" + err.Error())
	}

	if _, err := d.SendCommand("environment more false"); err != nil {
		return errors.New("failed to disable pagination" + err.Error())
	}

	return nil
}
