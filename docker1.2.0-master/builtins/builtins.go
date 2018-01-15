package builtins

import (
	"runtime"

	"github.com/CliffYuan/docker1.2.0/api"
	apiserver "github.com/CliffYuan/docker1.2.0/api/server"
	"github.com/CliffYuan/docker1.2.0/daemon/networkdriver/bridge"
	"github.com/CliffYuan/docker1.2.0/dockerversion"
	"github.com/CliffYuan/docker1.2.0/engine"
	"github.com/CliffYuan/docker1.2.0/events"
	"github.com/CliffYuan/docker1.2.0/pkg/parsers/kernel"
	"github.com/CliffYuan/docker1.2.0/registry"
)
//加载builtins主要是向engine注册多个handler(以便在执行相应的任务时，运行指定的handler)
func Register(eng *engine.Engine) error {
	//向engine注册网络初始化方法
    if err := daemon(eng); err != nil {
		return err
	}
    //向engine注册API服务处理方法(serverapi,acceptconnections)
	if err := remote(eng); err != nil {
		return err
	}
    //注册events事件处理方法(查看docker内部的event信息)
	if err := events.New().Install(eng); err != nil {
		return err
	}
    //注册版本处理方法
	if err := eng.Register("version", dockerVersion); err != nil {
		return err
	}
    //注册registry处理方法
	return registry.NewService().Install(eng)
}

// remote: a RESTful api for cross-docker communication
func remote(eng *engine.Engine) error {
	if err := eng.Register("serveapi", apiserver.ServeApi); err != nil {
		return err
	}
	return eng.Register("acceptconnections", apiserver.AcceptConnections)
}

// daemon: a default execution and storage backend for Docker on Linux,
// with the following underlying components:
//
// * Pluggable storage drivers including aufs, vfs, lvm and btrfs.
// * Pluggable execution drivers including lxc and chroot.
//
// In practice `daemon` still includes most core Docker components, including:
//
// * The reference registry client implementation
// * Image management
// * The build facility
// * Logging
//
// These components should be broken off into plugins of their own.
//
func daemon(eng *engine.Engine) error {
	return eng.Register("init_networkdriver", bridge.InitDriver)
}

// builtins jobs independent of any subsystem
func dockerVersion(job *engine.Job) engine.Status {
	v := &engine.Env{}
	v.SetJson("Version", dockerversion.VERSION)
	v.SetJson("ApiVersion", api.APIVERSION)
	v.Set("GitCommit", dockerversion.GITCOMMIT)
	v.Set("GoVersion", runtime.Version())
	v.Set("Os", runtime.GOOS)
	v.Set("Arch", runtime.GOARCH)
	if kernelVersion, err := kernel.GetKernelVersion(); err == nil {
		v.Set("KernelVersion", kernelVersion.String())
	}
	if _, err := v.WriteTo(job.Stdout); err != nil {
		return job.Error(err)
	}
	return engine.StatusOK
}
