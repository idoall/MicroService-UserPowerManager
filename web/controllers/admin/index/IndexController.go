package index

import (
	"bytes"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/columns"
	"github.com/idoall/TokenExchangeCommon/commonutils"
	"gitlab.mshk.top/TokenExchange/tokenexchangemodels/models"
)

var BaseURL = "in"
var (
	TemplageBaseURL = "index"
	baseTitle       = "首页"
)

// IndexController struct
type IndexController struct {
	admin.AdminBaseController
}

// IframeDefault 框架内的首页
func (e *IndexController) IframeDefault() {
	e.Data["title"] = fmt.Sprintf("%s", baseTitle)

	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts(nil)
	e.AppendCustomStyles(nil)

	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/iframe_index.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// Get 获取 Admin 默认带 Iframe 首页信息
func (e *IndexController) Get() {

	var result models.Result

	userID, userName, err := e.GetCurrentUser()
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["UserName"] = userName
	e.Data["UserID"] = userID
	e.Data["URL_IframeIndex"] = fmt.Sprintf("%s/%s/iframe_index", versionAdminURL, TemplageBaseURL)
	e.Data["LoginOutURL"] = beego.AppConfig.String("WebSite::URL_Logout")

	e.SetMortStype()
	e.SetMortScript()
	e.appendCustomScripts()
	e.TplName = fmt.Sprintf("%s/%s/index.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// GetAdminMenuHtml 返回管理员左侧菜单
func (e *IndexController) GetAdminMenuHTML(userID int64) string {
	columnList, err := new(columns.ColumnsController).GetTreeStruct()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	noNodesTemplate := `
	<li>
		<a class="J_menuItem" href="%s"><i class="%s"></i> <span class="nav-label">%s</span></a>
	</li>`
	NodesTemplate := `
	<li>
		<a class="J_menuItem" href="%s"><i class="%s"></i> <span class="nav-label">%s</span><span class="fa arrow"></span></a>
		%s
	</li>`
	NodesSecondULTemplate := `
	<ul class="nav nav-second-level">
	%s
	</ul>`
	NodesThirdULTemplate := `
	<ul class="nav nav-third-level">
	%s
	</ul>`
	NodesLINoIconTemplate := `
	<li>
		<a class="J_menuItem" href="%s">%s <span class="fa arrow"></span></a>
		%s
	</li>
	`
	noNodesLINoIconTemplate := `
	<li>
		<a class="J_menuItem" href="%s">%s</a>
	</li>
	`

	//Buffer是一个实现了读写方法的可变大小的字节缓冲
	var buffer bytes.Buffer
	for _, v := range columnList {
		//如果没有一级菜单权限 continue
		if !e.HasPermissions(userID, v.ID) {
			continue
		}

		// 如果没有二级菜单，直接展示一级菜单
		if v.Nodes == nil {
			//判断是否在首页显示
			if v.IsShowNav {
				buffer.WriteString(fmt.Sprintf(noNodesTemplate, v.URL, v.CssIcon, v.Name))
			}
		} else {
			var bufferSecond bytes.Buffer
			isHaveSvChildNodes := false //是否真的有下一级菜单
			for _, sv := range v.Nodes {

				//如果没有二级菜单权限 continue
				if !e.HasPermissions(userID, sv.ID) {
					continue
				}

				//如果有二级菜单，并且二级菜单是有显示的，在一级菜单上显示 左侧箭头
				if sv.IsShowNav {
					isHaveSvChildNodes = true
				}
				// 如果没有三级菜单，直接展示二级菜单
				if sv.Nodes == nil {
					//判断是否在首页显示
					if sv.IsShowNav {
						sHTML := fmt.Sprintf(NodesSecondULTemplate, fmt.Sprintf(noNodesLINoIconTemplate, sv.URL, sv.Name))
						bufferSecond.WriteString(sHTML)
					}
				} else {
					var bufferThird bytes.Buffer
					isHaveTvChildNodes := false //是否真的有下一级菜单
					for _, tv := range sv.Nodes {
						//如果没有三级菜单权限 continue
						if !e.HasPermissions(userID, tv.ID) {
							continue
						}
						if tv.IsShowNav {

							isHaveTvChildNodes = true
							tHTML := fmt.Sprintf(noNodesLINoIconTemplate, tv.URL, tv.Name)
							bufferThird.WriteString(fmt.Sprintf(NodesThirdULTemplate, tHTML))
						}
					}

					//是否有下级菜单，如果有显示最右侧的箭头
					if isHaveTvChildNodes {
						sHTML := fmt.Sprintf(NodesSecondULTemplate, fmt.Sprintf(NodesLINoIconTemplate, sv.URL, sv.Name, bufferThird.String()))
						bufferSecond.WriteString(sHTML)
					} else {
						sHTML := fmt.Sprintf(NodesSecondULTemplate, fmt.Sprintf(noNodesLINoIconTemplate, sv.URL, sv.Name))
						bufferSecond.WriteString(sHTML)
					}
				}

			}

			//是否有下级菜单，如果有显示最右侧的箭头
			if isHaveSvChildNodes {
				liStr := fmt.Sprintf(NodesTemplate, v.URL, v.CssIcon, v.Name, bufferSecond.String())
				buffer.WriteString(liStr)
			} else {
				buffer.WriteString(fmt.Sprintf(noNodesTemplate, v.URL, v.CssIcon, v.Name))
			}
		}
	}

	//替换掉下拉列表中的特殊字符
	replaceHTML := func(str string) string {
		str = commonutils.ReplaceString(str, "✚", "", -1)
		str = commonutils.ReplaceString(str, "┊", "", -1)
		str = commonutils.ReplaceString(str, "├", "", -1)
		return str
	}
	return replaceHTML(buffer.String())
}

//注册自定义脚本
func (e *IndexController) appendCustomScripts() {
	e.Data["customScripts"] = []string{
		"/static/js/hplus/plugins/metisMenu/jquery.metisMenu.js",
		"/static/js/hplus/plugins/slimscroll/jquery.slimscroll.min.js",
		"/static/js/hplus/plugins/layer/layer.min.js",
		"/static/js/hplus/hplus.js?v=4.1.0",         //自定义js
		"/static/js/hplus/contabs.js?1",             //自定义js
		"/static/js/hplus/plugins/pace/pace.min.js", //第三方插件
	}
}
