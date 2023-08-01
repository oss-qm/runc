// +build linux,no_systemd

package systemd

import (
	"errors"
	"github.com/opencontainers/runc/libcontainer/configs"
	"github.com/opencontainers/runc/libcontainer/cgroups"
)

func IsRunningSystemd() bool {
	return false
}

func NewUnifiedManager(config *configs.Cgroup, path string) (cgroups.Manager, error) {
    return nil, errors.New("no lennartix")
}

func NewLegacyManager(cg *configs.Cgroup, paths map[string]string) (cgroups.Manager, error) {
    return nil, errors.New("no lennartix")
}
