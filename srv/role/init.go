package main

import (
	"fmt"
	"net/url"
	"os"
	"runtime"

	"github.com/astaxie/beego/orm"
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/idoall/MicroService-UserPowerManager/log4"
	"github.com/idoall/MicroService-UserPowerManager/srv/role/v1/handler"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

// main init
func init() {

	// 通用变量
	var err error

	// 建议程序开启多核支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Logger
	logName := fmt.Sprintf("access-%s.log", "userpowermanager")
	inner.Mlogger = log4.NewFileLogger(logName)

	// 注册 mysql 驱动
	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		inner.Mlogger.Fatal(fmt.Sprintf("main init 注册数据库时出现错误 Error:%s", err.Error()))
		// fmt.Fprintln(os.Stdout, gocolorize.NewColor("red").Paint("main init 注册数据库时出现错误 Err:", err))
		return
	}

	dbConnStr := GetDBConnStr()
	// 注册数据库
	if err = orm.RegisterDataBase("default", "mysql", dbConnStr, 0, 1000); err != nil {
		inner.Mlogger.Fatalf(fmt.Sprintf("orm.RegisterDataBase dbStr:%s Error:", dbConnStr) + err.Error())
		return
	} else if utils.RunMode == "dev" {
		inner.Mlogger.Info("数据库注册成功")
		inner.Mlogger.Infof("DBConStr:%s", dbConnStr)
	}

	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := xormadapter.NewAdapter("mysql", dbConnStr)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	handler.RoleS = casbin.NewEnforcer("conf/rbac_model.conf", a)
	// fmt.Println(RoleS.Enforce("alice", "data1", "read"))
	// Load the policy from DB.
	handler.RoleS.LoadPolicy()

}

// 获取数据库连接字符串
// 配置读取优先级 环境变量->配置文件
func GetDBConnStr() string {
	host := utils.TConfig.String("Database::Host")
	if os.Getenv("DATABASE_SERVER_HOST") != "" {
		host = os.Getenv("DATABASE_SERVER_HOST")
	}

	port := utils.TConfig.String("Database::Port")
	if os.Getenv("DATABASE_SERVER_PORT") != "" {
		port = os.Getenv("DATABASE_SERVER_PORT")
	}

	dbName := utils.TConfig.String("Database::DBName")
	if os.Getenv("DATABASE_SERVER_DBNAME") != "" {
		dbName = os.Getenv("DATABASE_SERVER_DBNAME")
	}

	userName := utils.TConfig.String("Database::UserName")
	if os.Getenv("DATABASE_SERVER_USERNAME") != "" {
		userName = os.Getenv("DATABASE_SERVER_USERNAME")
	}

	password := utils.TConfig.String("Database::PassWord")
	if os.Getenv("DATABASE_SERVER_PASSWORD") != "" {
		password = os.Getenv("DATABASE_SERVER_PASSWORD")
	}

	charset := utils.TConfig.String("Database::Charset")
	if os.Getenv("DATABASE_SERVER_CHARSET") != "" {
		charset = os.Getenv("DATABASE_SERVER_CHARSET")
	}

	// loc := os.Getenv("DATABASE_SERVER_LOC")
	// if loc == "" {
	// 	loc = DATABASE_SERVER_LOC
	// }
	//
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		userName,
		password,
		host,
		port,
		dbName,
	)
	dbConnURLQuery := url.Values{}
	dbConnURLQuery.Set("charset", charset)
	dbConnURLQuery.Set("loc", "Asia/Shanghai")
	dbConnStr += "?" + dbConnURLQuery.Encode()
	return dbConnStr
}
