package main

import (
	"fmt"
	"runtime"

	"github.com/idoall/MicroService-UserPowerManager/log4"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

func init() {

	// 建议程序开启多核支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Logger
	logName := fmt.Sprintf("access-%s.log", "userpowermanager")
	inner.Mlogger = log4.NewFileLogger(logName)
}
