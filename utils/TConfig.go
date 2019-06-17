package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

var (
	// TConfig is the default config for Application
	TConfig config.Configer
	// AppPath is the absolute path to the app
	AppPath string
	// DataStrategyPath is the absolute path to the app
	DataStrategyPath string
	// appConfigPath is the path to the config files
	appConfigPath string
	// Verbose 是否打印日志信息
	Verbose bool
	// RunMode   dev | prod
	RunMode string

	// AdminVerifyLogin 是否开始管理员验证，主要后台用于调试用
	AdminVerifyLogin bool
)

func init() {
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	DataStrategyPath = filepath.Join(workPath, "data", "Strategy")

	var filename = "app.conf"
	appConfigPath = filepath.Join(workPath, "conf", filename)
	if !commonutils.PathExists(appConfigPath) {
		panic(fmt.Errorf("not find ./conf/app.conf"))
	}

	TConfig, err = config.NewConfig("ini", appConfigPath)
	if err != nil {
		panic(err)
	}

	RunMode = TConfig.String("default::runmode")
	if RunMode != "dev" && RunMode != "prod" {
		panic("错误的runmode参数， [dev | prod]")
	}

	if os.Getenv("MICROSERVICE_HOSTPORT") != "" {
		inner.MicroServiceHostProt = os.Getenv("MICROSERVICE_HOSTPORT")
	} else {
		inner.MicroServiceHostProt = "http://localhost:8080"
	}

	if AdminVerifyLogin, err = TConfig.Bool("WebSite::AdminVerifyLogin"); err != nil {
		panic("Conf 读取不到 AdminVerifyLogin")
	}
}
