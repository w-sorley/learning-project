package execdrivers

import (
	"fmt"
	"github.com/CliffYuan/docker1.2.0/daemon/execdriver"
	"github.com/CliffYuan/docker1.2.0/daemon/execdriver/lxc"
	"github.com/CliffYuan/docker1.2.0/daemon/execdriver/native"
	"github.com/CliffYuan/docker1.2.0/pkg/sysinfo"
	"path"
)

func NewDriver(name, root, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
	switch name {
	case "lxc":
		// we want to give the lxc driver the full docker root because it needs
		// to access and write config and template files in /var/lib/docker/containers/*
		// to be backwards compatible
		return lxc.NewDriver(root, initPath, sysInfo.AppArmor)
	case "native":
		return native.NewDriver(path.Join(root, "execdriver", "native"), initPath)
	}
	return nil, fmt.Errorf("unknown exec driver %s", name)
}
