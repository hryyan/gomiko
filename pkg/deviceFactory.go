package gomiko

import (
	"strings"

	"github.com/hryyan/gomiko/pkg/connections"
	"github.com/hryyan/gomiko/pkg/lib/arista"
	"github.com/hryyan/gomiko/pkg/lib/cisco"
	"github.com/hryyan/gomiko/pkg/lib/juniper"
	"github.com/hryyan/gomiko/pkg/lib/mikrotik"
	"github.com/hryyan/gomiko/pkg/lib/sros"
	"github.com/hryyan/gomiko/pkg/lib/centec"
	"github.com/hryyan/gomiko/pkg/types"
	"github.com/pkg/errors"
)

func NewDevice(Host string, Username string, Password string, DeviceType string, Port uint8, Options ...DeviceOption) (types.Device, error) {
	var device types.Device

	//for Mikrotik you need to append +ct200w to username
	if strings.Contains(DeviceType, "mikrotik") {
		Username += "+ct200w"
	}

	//create connection
	connection, err := connections.NewConnection(Host, Username, Password, "ssh", Port)
	if err != nil {
		return nil, err
	}

	//create the Device
	if strings.Contains(DeviceType, "cisco") {
		device, err = cisco.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		device, err = arista.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		device, err = juniper.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
		device, err = mikrotik.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "nokia_sros") {
		device, err = sros.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "centec") {
		device, err = centec.NewDevice(connection, DeviceType)
	} else {
		return nil, errors.New("DeviceType not supported: " + DeviceType)
	}
	if err != nil {
		return nil, err
	}

	// running Options Functions.
	for _, option := range Options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}

	return device, nil

}
