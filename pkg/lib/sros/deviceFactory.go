package sros

import (
	"github.com/hryyan/gomiko/pkg/connections"
	"github.com/hryyan/gomiko/pkg/driver"
	"github.com/hryyan/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\n")
	return &SROSDevice{
		Prompt: "",
		Driver: devDriver,
	}, nil
}
