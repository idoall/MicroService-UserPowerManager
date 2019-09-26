package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/idoall/MicroService-UserPowerManager/common"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/log4"
	"github.com/idoall/MicroService-UserPowerManager/utils/orm"
)

// main init
func init() {

	common.OutBanner()
	common.OutVersion()

	// HandleInterrupt 捕获 goroutine 的退出信号
	inner.HandleInterrupt()

	// Init Logger
	logName := fmt.Sprintf("access-%s.log", inner.LOGFILENAME)
	inner.Mlogger = log4.NewFileLogger(logName)

	// 建议程序开启多核支持
	common.AdjustGoMaxProcs()

	// 使用 XROM
	orm.InitXorm()

}
