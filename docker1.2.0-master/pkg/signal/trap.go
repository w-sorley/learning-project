package signal

import (
	"log"
	"os"
	gosignal "os/signal"
	"sync/atomic"
	"syscall"
)

// Trap sets up a simplified signal "trap", appropriate for common
// behavior expected from a vanilla unix command-line tool in general
// (and the Docker engine in particular).
//
// * If SIGINT or SIGTERM are received, `cleanup` is called, then the process is terminated.
// * If SIGINT or SIGTERM are repeated 3 times before cleanup is complete, then cleanup is
// skipped and the process terminated directly.
// * If "DEBUG" is set in the environment, SIGQUIT causes an exit without cleanup.
//
//设置信号捕获
func Trap(cleanup func()) {
    //　创建一个channel，用于发送信号通知
	c := make(chan os.Signal, 1)
    //　定义sinal数组变量，并设置初始值
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	if os.Getenv("DEBUG") == "" {
		signals = append(signals, syscall.SIGQUIT)
	}
    //通过Notify函数将接收到的信号传递给刚才创建的channel（只有声明的信号类型才会传递，其余会被忽略）
	gosignal.Notify(c, signals...)
	//创建一个goroutine用于处理信号
    go func() {
		interruptCount := uint32(0)
		for sig := range c {
			go func(sig os.Signal) {
				log.Printf("Received signal '%v', starting shutdown of docker...\n", sig)
			//根据信号的类型，执行对应的处理逻辑
                switch sig {
				case os.Interrupt, syscall.SIGTERM:
					// If the user really wants to interrupt, let him do so.
					if atomic.LoadUint32(&interruptCount) < 3 {
						atomic.AddUint32(&interruptCount, 1)
						// Initiate the cleanup only once
						if atomic.LoadUint32(&interruptCount) == 1 {
							// Call cleanup handler
							cleanup()
							os.Exit(0)
						} else {
							return
						}{
					} else {
						log.Printf("Force shutdown of docker, interrupting cleanup\n")
					}
				case syscall.SIGQUIT:
				}
				os.Exit(128 + int(sig.(syscall.Signal)))
			}(sig)
		}
	}()
}
