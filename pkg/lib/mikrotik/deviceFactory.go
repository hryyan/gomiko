package mikrotik

import (
	"errors"

	"github.com/hryyan/gomiko/pkg/connections"
	"github.com/hryyan/gomiko/pkg/driver"
	"github.com/hryyan/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\r")

	switch DeviceType {
	case "mikrotik_routeros":
		return &MikroTikRouterOS{
			Driver:     devDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
