package admin

import (
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

var (

	// AdminVerifyLogin 是否开始管理员验证，主要后台用于调试用
	AdminVerifyLogin bool
)

func init() {
	var err error
	if AdminVerifyLogin, err = utils.TConfig.Bool("WebSite::AdminVerifyLogin"); err != nil {
		inner.Mlogger.Error("Conf 读取不到 AdminVerifyLogin")
	}
}
