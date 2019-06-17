package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// Result struct
type Result struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
	Msg   string `json:"msg"`
}

// AdminBaseController struct
type AdminBaseController struct {
	controllers.BaseController
}

var (
	AdminBaseRoterVersion = "v1"
	TemplageAdminBaseURL  = "admin"
)

// 返回版本号+admin拼接的url
func (e *AdminBaseController) GetVersionAdminBaseURL() string {
	return fmt.Sprintf("/%s/%s", AdminBaseRoterVersion, TemplageAdminBaseURL)
}

// AppendCustomStyles 注册自定义脚样式
func (e *AdminBaseController) AppendCustomStyles(list []string) {
	styleList := []string{
		"/static/css/hplus/iCheck/custom.css",
		"/static/css/hplus/chosen/chosen.css",
		"/static/css/awesome-bootstrap-checkbox/awesome-bootstrap-checkbox.css",
		"/static/css/hplus/sweetalert/sweetalert.css",
	}
	for _, v := range list {
		if !commonutils.StringDataContains(styleList, v) {
			styleList = append(styleList, v)
		}
	}
	e.Data["customStyles"] = styleList
}

// AppendCustomScripts 注册自定义脚本
func (e *AdminBaseController) AppendCustomScripts(list []string) {
	scriptList := []string{
		"/static/js/hplus/plugins/validate/jquery.validate.min.js",
		"/static/js/hplus/plugins/validate/messages_zh.min.js",
		// "/static/js/hplus/content.js?v=1.0.0",
		"/static/js/hplus/plugins/iCheck/1.0.2/icheck.min.js",
		// "/static/js/ianPager.js",
		"/static/js/hplus/plugins/sweetalert/2.1.0/sweetalert.min.js", //第三方插件
		// "/static/js/hplus/plugins/echarts/4.1.0/echarts.min.js",       //百度echarets
		//layer
		"/static/js/hplus/plugins/layer/layer.min.js",
	}
	for _, v := range list {
		if !commonutils.StringDataContains(scriptList, v) {
			scriptList = append(scriptList, v)
		}
	}
	e.Data["customScripts"] = scriptList
}

// GetCurrentUser 返回用户ID
func (e *AdminBaseController) GetCurrentUser() (int64, error) {
	tokenString := e.Ctx.Input.Cookie("mshk_token")
	if tokenString == "" {
		return int64(0), errors.New("mshk_token is nil")
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Token", tokenString)

	// 临时 Json 解析类
	responseJSON := struct {
		Vaild       int    `json:"vaild"`       // 是否验证通过
		UserID      string `json:"userid"`      //用户ID
		UserName    string `json:"username"`    //用户登录名
		TokenString string `json:"tokenstring"` //返回的 Token
	}{}
	// 发送 http 请求
	if err := request.Request.WebPOSTSendPayload("ServiceURL_User_ValidToken", bytes.NewBufferString(params.Encode()), &responseJSON); err != nil {
		inner.Mlogger.Error(err)
		// 转到登录
		return int64(0), err
	}

	return strconv.ParseInt(responseJSON.UserID, 10, 64)
}

// HasPermissions 是否具有权限
func (e *AdminBaseController) HasPermissions(userID, powerID int64) bool {
	var err error
	//如果没有打开验证，直接通过
	if !utils.AdminVerifyLogin {
		return true
	}
	var power bool

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Name", fmt.Sprintf("%d", userID))

	var responseRolsJSON []string

	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_GetRolesForUser", bytes.NewBufferString(params.Encode()), responseRolsJSON); err != nil {
		inner.Mlogger.Error("HasPermissions ServiceURL_Role_GetRolesForUser Error:" + err.Error())
		return false
	}

	powerIDString := fmt.Sprintf("%d", powerID)
	for _, v := range responseRolsJSON {

		params = url.Values{}
		params.Set("User", v)

		var responsePermissionsJSON []map[string][]string

		if err = request.Request.WebGETSendPayload("ServiceURL_Role_GetPermissionsForUser", params, &responseJSON); err != nil {
			inner.Mlogger.Error("HasPermissions ServiceURL_Role_GetPermissionsForUser Error:" + err.Error())
			return false
		}

		for _, ov := responsePermissionsJSON{
			if ov["Two"][1] == powerIDString{
				return true
			}
		}
	}
	return power
}
