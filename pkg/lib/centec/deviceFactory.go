package centec

import (
	"github.com/hryyan/gomiko/pkg/connections"
	"github.com/hryyan/gomiko/pkg/driver"
	"github.com/hryyan/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDevice := driver.NewDriver(connection, "\n")
	return &CnosDevice{
		Prompt: "",
		Driver: devDevice,
	}, nil
}
