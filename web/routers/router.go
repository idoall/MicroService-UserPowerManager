// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/siteauth"

	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/columns"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/index"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/users"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/usersgroup"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {

	// 注册后台管理 - 首页模块
	routerIndexModel := &index.IndexController{}
	// 注册后台管理 - 用户模块
	routerUserModel := &users.UsersController{}
	// 注册后台管理 - 栏目模块
	routerColumnsModel := &columns.ColumnsController{}
	// 注册后台管理 - 用户组模块
	routerUsersGroupModel := &usersgroup.UsersGroupController{}
	// 注册登录、退出模块
	routerSiteAuthModel := &siteauth.SiteAuthController{}

	beego.Get("/ping", func(ctx *context.Context) {
		ctx.Output.Body([]byte("pong"))
	})
	ns := beego.NewNamespace(fmt.Sprintf("/%s", admin.AdminBaseRoterVersion),

		// beego.NSCond(func(ctx *context.Context) bool {
		// 	if ctx.Input.Domain() == "api.beego.me" {
		// 		return true
		// 	}
		// 	return false
		// }),
		// beego.NSRouter("*",&controllers.BaseController{},"OPTIONS:Options"),)
		// 判断 UserAgent 不能为空
		beego.NSCond(func(ctx *context.Context) bool {
			if ua := ctx.Input.UserAgent(); ua != "" {
				return true
			}
			return false
		}),
		// admin 后台
		beego.NSNamespace(fmt.Sprintf("/%s", admin.TemplageAdminBaseURL),

			beego.NSBefore(filterAdminUserLoginGateWay), //权限验证

			// 首页
			beego.NSRouter("/", routerIndexModel),
			beego.NSNamespace(fmt.Sprintf("/%s", index.TemplageBaseURL),
				// 首页，默认调用 Get 方法
				beego.NSRouter("/", routerIndexModel),
			),

			// 用户管理
			beego.NSNamespace(fmt.Sprintf("/%s", users.TemplageBaseURL),
				// 首页，默认调用 Get 方法
				beego.NSRouter("/", routerUserModel),
				// 获取JSON列表
				beego.NSRouter("/GetListJSON", routerUserModel, "*:GetListJSON"),
				// 添加用户
				beego.NSRouter("/add", routerUserModel, "*:Add"),
				// 添加用户 - 保存
				beego.NSRouter("/addsave", routerUserModel, "*:AddSave"),
				// 修改用户
				beego.NSRouter("/update", routerUserModel, "*:Update"),
				// 修改用户 - 保存
				beego.NSRouter("/updatesave", routerUserModel, "*:UpdateSave"),
				// 删除用户
				beego.NSRouter("/delete", routerUserModel, "*:Delete"),
				// 批量删除用户
				beego.NSRouter("/batchdelete", routerUserModel, "*:BatchDelete"),
			),
			// 栏目管理
			beego.NSNamespace(fmt.Sprintf("/%s", columns.TemplageBaseURL),
				// 首页，默认调用 Get 方法
				beego.NSRouter("/", routerColumnsModel),
				// 获取JSON列表
				beego.NSRouter("/GetTreeViewJSON", routerColumnsModel, "*:GetTreeViewJSON"),
				// 添加
				beego.NSRouter("/add", routerColumnsModel, "*:Add"),
				// 添加 - 保存
				beego.NSRouter("/addsave", routerColumnsModel, "*:AddSave"),
				// 修改
				beego.NSRouter("/update", routerColumnsModel, "*:Update"),
				// 修改 - 保存
				beego.NSRouter("/updatesave", routerColumnsModel, "*:UpdateSave"),
				// 批量删除
				beego.NSRouter("/batchdelete", routerColumnsModel, "*:BatchDelete"),
			),
			// 用户组管理
			beego.NSNamespace(fmt.Sprintf("/%s", usersgroup.TemplageBaseURL),
				// 首页，默认调用 Get 方法
				beego.NSRouter("/", routerUsersGroupModel),
				// 获取JSON列表
				beego.NSRouter("/GetListJSON", routerUsersGroupModel, "*:GetListJSON"),
				// 添加
				beego.NSRouter("/add", routerUsersGroupModel, "*:Add"),
				// 添加 - 保存
				beego.NSRouter("/addsave", routerUsersGroupModel, "*:AddSave"),
				// 修改
				beego.NSRouter("/update", routerUsersGroupModel, "*:Update"),
				// 修改 - 保存
				beego.NSRouter("/updatesave", routerUsersGroupModel, "*:UpdateSave"),
				// 删除用户
				beego.NSRouter("/delete", routerUsersGroupModel, "*:Delete"),
				// 批量删除
				beego.NSRouter("/batchdelete", routerUsersGroupModel, "*:BatchDelete"),
				// 获取栏目 JSON 列表
				beego.NSRouter("/GetColumnPowerTreeViewJSON", routerUsersGroupModel, "*:GetColumnPowerTreeViewJSON"),
				// 获取栏目权限
				beego.NSRouter("/ColumnPower", routerUsersGroupModel, "*:ColumnPower"),
				// 更改权限
				beego.NSRouter("/ColumnPowerSave", routerUsersGroupModel, "*:ColumnPowerSave"),
			),
		),

		// 登录
		beego.NSNamespace(fmt.Sprintf("/%s", siteauth.TemplageBaseURL),
			// 首页，登录
			beego.NSRouter("/", routerSiteAuthModel, "*:Login"),
			// 退出登录
			beego.NSRouter("/logout", routerSiteAuthModel, "*:Logout"),
			// 检测用懚是否可以登录
			beego.NSRouter("/checklogin", routerSiteAuthModel, "*:CheckLogin"),
		),
	)
	beego.AddNamespace(ns)
}

// 访问 admin 模块下做的权限验证
func filterAdminUserLoginGateWay(ctx *context.Context) {

	// 获取 cookie 中的 token
	tokenString := ctx.Input.Cookie("mshk_token")
	if tokenString == "" {
		// 转到登录
		ctx.Redirect(302, fmt.Sprintf("/%s%s?Referer=", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_Login"))+url.QueryEscape(ctx.Request.RequestURI))
		return
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
		ctx.Redirect(302, fmt.Sprintf("/%s%s?Referer=", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_Login"))+url.QueryEscape(ctx.Request.RequestURI))
		return
	} else {
		if commonutils.StringContains(utils.RunMode, "dev") {
			inner.Mlogger.Info(fmt.Sprintf("用户 %s %s Token验证成功", responseJSON.UserName, responseJSON.UserID))
		}

		// 写入 cookie,10分钟后过期
		ctx.Output.Cookie("mshk_token", responseJSON.TokenString, 60*10)
	}
}
