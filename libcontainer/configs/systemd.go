package configs

import (
	systemdDbus "github.com/coreos/go-systemd/v22/dbus"
)

type SdProperty   = systemdDbus.Property
type SdProperties = []systemdDbus.Property
