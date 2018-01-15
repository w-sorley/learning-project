// +build daemon

package main

import (
	"log"

	"github.com/CliffYuan/docker1.2.0/builtins"
	"github.com/CliffYuan/docker1.2.0/daemon"
	_ "github.com/CliffYuan/docker1.2.0/daemon/execdriver/lxc"
	_ "github.com/CliffYuan/docker1.2.0/daemon/execdriver/native"
	"github.com/CliffYuan/docker1.2.0/dockerversion"
	"github.com/CliffYuan/docker1.2.0/engine"
	flag "github.com/CliffYuan/docker1.2.0/pkg/mflag"
	"github.com/CliffYuan/docker1.2.0/pkg/signal"
)

const CanDaemon = true
//声明daemon配置信息变量，并初始化赋值
var (
	daemonCfg = &daemon.Config{}
)
func init() {
	daemonCfg.InstallFlags()
}

//根据用户命令(-d)启动一个docker daemon守护进程
func mainDaemon() {
      //判断剩余的参数是否为0,若传入多余参数则用法有误(启动Docker daemon时无需多余参数)
	if flag.NArg() != 0 {
		flag.Usage()
		return
	}
        //创建Engine(运行引擎，通过job管理所有运行中的任务)
	eng := engine.New()
       //设置engine的信号捕获(当遇到SIGINT,SIGTERM,SIGQUIT信号运行相应函数，执行善后并退出,shutdown函数主要时执行善后清理工作)，daemon作为linux后台进程应该具有处理信号的能力，使得用户可以以向其发送信号的形式管理其运行
	signal.Trap(eng.Shutdown)
	// Load builtins加载builtins（与容器无关，与docker的运行时信息有关，注册相关命令对应的handler）
	if err := builtins.Register(eng); err != nil {
		log.Fatal(err)
	}

	// load the daemon in the background so we can immediately start
	// the http api so that connections don't fail while the daemon
	// is booting
    //使用goroutine加载daemon对象并运行docker server
	go func() {
        //创建daemon对象（核心，）
		d, err := daemon.NewDaemon(daemonCfg, eng)
		if err != nil {
			log.Fatal(err)
		}
        //通过daemon对象，向engine注册handler
		if err := d.Install(eng); err != nil {
			log.Fatal(err)
		}
		// after the daemon is done setting up we can tell the api to start
		// accepting connections
        //运行对应的job,通知init守护进程，使得docker daemon开始接受请求
		if err := eng.Job("acceptconnections").Run(); err != nil {
			log.Fatal(err)
		}
	}()
    //创建完成：打印版本和驱动信息
	// TODO actually have a resolved graphdriver to show?
	log.Printf("docker daemon: %s %s; execdriver: %s; graphdriver: %s",
		dockerversion.VERSION,
		dockerversion.GITCOMMIT,
		daemonCfg.ExecDriver,
		daemonCfg.GraphDriver,
	)
　　//随后立即创建并运行serverapi的job,提供(daemon对client的)API服务
	// Serve api
	job := eng.Job("serveapi", flHosts...)
	job.SetenvBool("Logging", true)
	job.SetenvBool("EnableCors", *flEnableCors)
	job.Setenv("Version", dockerversion.VERSION)
	job.Setenv("SocketGroup", *flSocketGroup)

	job.SetenvBool("Tls", *flTls)
	job.SetenvBool("TlsVerify", *flTlsVerify)
	job.Setenv("TlsCa", *flCa)
	job.Setenv("TlsCert", *flCert)
	job.Setenv("TlsKey", *flKey)｛
	job.SetenvBool("BufferRequests", true)
	if err := job.Run(); err != nil {
		log.Fatal(err)
	}
}
