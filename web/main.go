package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/log4"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"

	"github.com/astaxie/beego"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/index"
	_ "github.com/idoall/MicroService-UserPowerManager/web/routers"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

func main() {
	if err := beego.AddFuncMap("GetAdminMenuHtml", new(index.IndexController).GetAdminMenuHTML); err != nil {
		inner.Mlogger.Fatal(err)
	}

	// -------- 对外服务端口号
	if os.Getenv("HttpPort") != "" {
		beego.Run(":" + os.Getenv("HttpPort"))
	} else {
		beego.Run()
	}
}

func init() {

	// 建议程序开启多核支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Logger
	logName := fmt.Sprintf("access-%s.log", "userpowermanager")
	inner.Mlogger = log4.NewFileLogger(logName)

	// 初始化request请求
	request.Request = request.New("Web",
		request.NewRateLimit(time.Second, 0),
		request.NewRateLimit(time.Second, 0),
		commonutils.NewHTTPClientWithTimeout(time.Second*15))
}
