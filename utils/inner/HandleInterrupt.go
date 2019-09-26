package inner

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// HandleInterrupt 捕获 goroutine 的退出信号
func HandleInterrupt() {
	c := make(chan os.Signal, 1)

	// 使用 signal.Notify 注册要接收的信号
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		if sig.String() == "interrupt" {
			fmt.Printf("\n退出类别 %v, 手动退出. \n", sig)
			// os.Exit(0)
		} else {
			fmt.Printf("Captured %v, shutdown requested. \n", sig)
		}
	}()
}
