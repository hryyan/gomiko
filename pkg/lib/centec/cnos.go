package centec

import "github.com/hryyan/gomiko/pkg/driver"
import "errors"
import "strings"

type CnosDevice struct {
	Driver driver.IDriver
	DeviceType string
	Prompt string
}

func (d *CnosDevice) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("(.+?)#", "#")
	if err != nil {
		return err
	}
	d.Prompt = prompt

	return d.sessionPreparation()
}

func (d *CnosDevice) Disconnect() {
	d.Driver.Disconnect()
}

func (d *CnosDevice) SendCommand(command string) (string, error) {
	return d.Driver.SendCommand(command, d.Prompt)
}

func (d *CnosDevice) SendConfigSet(configs []string) (string, error) {
	result := ""
	for _, config := range configs {
		if r, err := d.Driver.SendCommand(config, d.Prompt); err != nil {
			result += r
			return result, err
		}
	}
	return result, nil
}

func (d *CnosDevice) sessionPreparation() error {

	out, err := d.Driver.SendCommand("cd", d.Prompt)
	if err != nil {
		return errors.New("Failed to send cli command" + err.Error())
	}
	if !strings.Contains(out, "#") {
		return errors.New("Failed to send cli command" + err.Error())
	}

	return nil
}



