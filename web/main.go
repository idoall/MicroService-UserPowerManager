package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/idoall/MicroService-UserPowerManager/log4"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"

	"github.com/astaxie/beego"
	_ "github.com/idoall/MicroService-UserPowerManager/web/routers"
)

func main() {
	beego.Run()
}

func init() {

	// 建议程序开启多核支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Logger
	logName := fmt.Sprintf("access-%s.log", "userpowermanager")
	inner.Mlogger = log4.NewFileLogger(logName)

	// 初始化request请求
	request.Request = request.New("UserPowerManager",
		&http.Client{Timeout: (time.Second * 15)})
}
