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

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/siteauth"

	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/columns"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/users"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/usersgroup"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {

	// 注册用户模块
	routerUserModel := &users.UsersController{}
	// 注册栏目模块
	routerColumnsModel := &columns.ColumnsController{}
	// 注册用户组模块
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
			// beego.NSBefore(filterAdminUserLoginGateWay), //权限验证
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
			),
		),

		// 登录
		beego.NSNamespace(fmt.Sprintf("/%s", siteauth.TemplageBaseURL),
			// 首页，登录
			beego.NSRouter("/", routerSiteAuthModel, "*:Login"),
			// 退出登录
			beego.NSRouter("/loginout", routerSiteAuthModel, "*:LoginOut"),
			// 检测用懚是否可以登录
			beego.NSRouter("/checklogin", routerSiteAuthModel, "*:CheckLogin"),
		),
	)
	beego.AddNamespace(ns)
}

// 访问 admin 模块下做的权限验证
func filterAdminUserLoginGateWay(ctx *context.Context) {

	// 获取 cookie 中的 token
	tokenString := ctx.Input.Cookie("token")
	if tokenString == "" {
		// 转到登录
		ctx.Redirect(302, fmt.Sprintf("/%s%s?Referer=", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_Login"))+url.QueryEscape(ctx.Request.RequestURI))
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Token", tokenString)

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_User_ValidToken"))

	// 临时 Json解析类
	responseJson := struct {
		TokenString string `json:"tokenstring"`
	}{}
	// 发送 http 请求
	if err := request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		inner.Mlogger.Error(err)
		// 转到登录
		ctx.Redirect(302, fmt.Sprintf("/%s%s?Referer=", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_Login"))+url.QueryEscape(ctx.Request.RequestURI))
		return
	} else {
		// 写入 cookie,10分钟后过期
		ctx.Output.Cookie("token", responseJson.TokenString, 60*10)
	}
}
