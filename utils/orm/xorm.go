package orm

import (
	"fmt"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

var (
	// XROM public
	XROM *xorm.EngineGroup
	// ConnentionString 数据库连接字符串
	ConnentionString string
)

// InitXorm 初始化 xorm
func InitXorm() {
	var err error
	var engine *xorm.Engine

	// 获取数据库连接字符串
	ConnentionString = getDBConnStr()

	if engine, err = xorm.NewEngine("mysql", ConnentionString); err != nil {
		inner.Mlogger.Fatal("xorm 注册数据库时出现错误 Error:" + err.Error())
	} else {

		if err = engine.Ping(); err != nil {
			inner.Mlogger.Fatal("xorm ping db Error:" + err.Error())
		}

		//负载均衡策略:(特性自行百度)
		// 1.xorm.RandomPolicy()随机访问负载均衡,
		// 2.xorm.WeightRandomPolicy([]int{2, 3,4})权重随机负载均衡
		// 3.xorm.RoundRobinPolicy() 轮询访问负载均衡
		// 4.xorm.WeightRoundRobinPolicy([]int{2, 3,4}) 权重轮训负载均衡
		// 5.xorm.LeastConnPolicy()最小连接数负载均衡
		if XROM, err = xorm.NewEngineGroup(engine, []*xorm.Engine{engine}, xorm.RandomPolicy()); err != nil {
			inner.Mlogger.Fatal("xorm.NewEngineGroup error:" + err.Error())
		}
		XROM.SetMaxIdleConns(0)
		XROM.SetMaxOpenConns(1000)
	}

	inner.Mlogger.Info("ConnentionString:" + ConnentionString)
}

// getDBConnStr 获取数据库连接字符串
// 配置读取优先级 环境变量->配置文件
func getDBConnStr() string {
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
