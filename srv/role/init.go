package main

import (
	"fmt"

	// beegoormadapter "github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/idoall/MicroService-UserPowerManager/common"
	"github.com/idoall/MicroService-UserPowerManager/srv/role/v1/handler"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/idoall/MicroService-UserPowerManager/utils/log4"
	"github.com/idoall/MicroService-UserPowerManager/utils/orm"
)

// main init
func init() {

	var err error

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

	//---- Beego ORM
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", dbConnStr, true)

	// handler.RoleS = casbin.NewEnforcer("conf/rbac_model.conf", a)
	// // fmt.Println(RoleS.Enforce("alice", "data1", "read"))
	// // Load the policy from DB.
	// handler.RoleS.LoadPolicy()

	//---- XORM
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, err := xormadapter.NewAdapter("mysql", orm.ConnentionString, true) // Your driver and data source.
	if err != nil {
		inner.Mlogger.Fatal("xormadapter.NewAdapter Error:" + err.Error())
	}
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	handler.RoleS, err = casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		inner.Mlogger.Fatal("casbin.NewEnforcer Error:" + err.Error())
	}
	// Load the policy from DB.
	if err = handler.RoleS.LoadPolicy(); err != nil {
		inner.Mlogger.Fatal("handler.RoleS.LoadPolicy Error:" + err.Error())
	}

	// 是否开启 Jaeger
	if enableJaeger, err := utils.TConfig.Bool("JaegerServices::Enable"); err != nil {
		inner.Mlogger.Fatalf("utils.TConfig Jaeger:%s", err)
	} else if enableJaeger {
		// 初始化 Jaeger 配置
		_, closer := jaeger.InitJaeger("UserPowerManager")
		defer closer.Close()
	}

}
