package usergroup

import (
	"fmt"

	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
)

// UsersGroupController Controller
type UsersGroupController struct {
	admin.AdminBaseController
}

var TemplageBaseURL = "usersgroup"
var baseTitle = "用户组"
var pageSizeDefault = 11

// GetListJSON Default Json
func (e *UsersGroupController) GetListJSON() {
	//get page
	// pageSize, _ := e.GetInt("pagesize", pageSizeDefault)
	// currentPageIndex, _ := e.GetInt("currentpage", 1)

	// jsonList := struct {
	// 	Rows       []*models.UsersGroup `json:"rows"`
	// 	Total      int64                `json:"total"`
	// 	ErrMessage string               `json:"errmsg"`
	// }{}

	// listCond := orm.NewCondition()
	// // userID := int64(6)
	// //可以查看所有用户列表
	// if !e.HasPermissions(e.GetUserID(), 22) {
	// 	listCond = listCond.And("id", -1)
	// }
	// list, totalcount, err := new(models.UserGroup).GetAll(listCond, pageSize, currentPageIndex, "-sort")
	// if err != nil {
	// 	jsonList.ErrMessage = err.Error()
	// } else {
	// 	jsonList.Total = totalcount
	// 	jsonList.Rows = list
	// }
	// e.Data["json"] = jsonList
	// e.ServeJSON()
}

// Get 首页
func (e *UsersGroupController) Get() {

	//set Data
	versionAdminURL := e.GetVersionAdminURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["URL_Add"] = fmt.Sprintf("/%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["URL_Update"] = fmt.Sprintf("/%s/%s/update?id=", versionAdminURL, TemplageBaseURL)
	e.Data["URL_Del"] = fmt.Sprintf("/%s/%s/del?id=", versionAdminURL, TemplageBaseURL)
	e.Data["URL_ColumnPower"] = fmt.Sprintf("/%s/%s/ColumnPower?id=", versionAdminURL, TemplageBaseURL)
	e.Data["URL_BatchDelete"] = fmt.Sprintf("/%s/%s/batchdel", versionAdminURL, TemplageBaseURL)
	e.Data["URL_JsonListUrl"] = fmt.Sprintf("/%s/%s/GetListJSON", versionAdminURL, TemplageBaseURL)
	// e.Data["Power_ViewUserGroupPowerBtn"] = e.HasPermissions(e.GetUserID(), 19) //显示配置权限按钮
	// e.Data["Power_ViewDelBtn"] = e.HasPermissions(e.GetUserID(), 25)            //显示删除和批量删除按钮

	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		//Bootstrap table
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table-mobile.min.js",
		"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
	})
	e.AppendCustomStyles([]string{
		//Bootstrap table
		"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
	})
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/index.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// Add 添加
func (e *UsersGroupController) Add() {

	// e.Data["title"] = fmt.Sprintf("添加%s", baseTitle)
	// e.Data["AddSaveUrl"] = fmt.Sprintf("/%s/addsave", BaseURL)

	// //公用设置，样式、脚本、layout
	// e.SetMortStype()
	// e.SetMortScript()
	// e.AppendCustomScripts([]string{"/static/js/admin/symbol_add.js"})
	// e.AppendCustomStyles(nil)
	// e.Layout = "admin/layout/layout.html"
	// e.LayoutSections = make(map[string]string)
	// e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	// e.TplName = BaseURL + "/add.html"
}

// AddSave 保存添加的交易配置
func (e *UsersGroupController) AddSave() {

	// var result models.Result

	// model := new(models.UserGroup)
	// model.Name = e.GetString("name")
	// model.Sort, _ = e.GetInt("sort", 0)
	// model.ParentID, _ = e.GetInt64("parintid", 0)
	// model.Note = e.GetString("note")
	// model.AddTime = time.Now().Format("2006-01-02 15:04:05")

	// _, err := model.Add(model)
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// } else {
	// 	e.Ctx.Redirect(302, fmt.Sprintf("/%s", BaseURL))
	// }
}

// Update 修改
func (e *UsersGroupController) Update() {
	// var result models.Result
	// id, _ := e.GetInt64("id", 0)
	// model, err := new(models.UserGroup).GetOne(id)

	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// //set Data
	// e.Data["Model"] = model
	// e.Data["title"] = fmt.Sprintf("修改%s", baseTitle)
	// e.Data["UpdateSaveUrl"] = fmt.Sprintf("/%s/updatesave", BaseURL)

	// //公用设置，样式、脚本、layout
	// e.SetMortStype()
	// e.SetMortScript()
	// e.AppendCustomScripts([]string{"/static/js/admin/symbol_add.js"})
	// e.AppendCustomStyles(nil)
	// e.Layout = "admin/layout/layout.html"
	// e.LayoutSections = make(map[string]string)
	// e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	// e.TplName = BaseURL + "/update.html"
}

// UpdateSave 保存修改
func (e *UsersGroupController) UpdateSave() {
	// var result models.Result
	// id, _ := e.GetInt64("id", 0)

	// model, err := new(models.UserGroup).GetOne(id)
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// model.Name = e.GetString("name")
	// model.Sort, _ = e.GetInt("sort", 0)
	// model.ParentID, _ = e.GetInt64("parintid", 0)
	// model.Note = e.GetString("note")

	// if num, err := model.Update(model); err == nil {
	// 	fmt.Println(num)
	// }
	// e.Ctx.Redirect(302, fmt.Sprintf("/%s", BaseURL))
}

// Delete 删除
func (e *UsersGroupController) Delete() {

	// id, _ := e.GetInt64("id", 0)
	// if id != 0 {
	// 	_, err := new(models.UserGroup).Delete(id)
	// 	if err != nil {
	// 		e.Data["json"] = err.Error()
	// 		e.ServeJSON()
	// 		return
	// 	}
	// }

	// e.Ctx.Redirect(302, fmt.Sprintf("/%s", BaseURL))
}

// BatchDelete 批量删除
func (e *UsersGroupController) BatchDelete() {

	// var result models.Result

	// if !e.HasPermissions(e.GetUserID(), 25) {
	// 	result.Code = -1
	// 	result.Msg = "没有删除权限"
	// 	e.Data["json"] = result
	// }

	// _, err := new(models.UserGroup).BatchDelete(strings.Split(e.GetString("ids"), ","))
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// } else {
	// 	result.Code = 0
	// 	e.Data["json"] = result
	// }
	// e.ServeJSON()
}

//-------------------权限配置

// ColumnPower 展示用户组-权限页面
func (e *UsersGroupController) ColumnPower() {
	// var result models.Result
	// id, _ := e.GetInt64("id", 0)

	// model, err := new(models.UserGroup).GetOne(id)
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// e.Data["Model"] = model
	// e.Data["title"] = fmt.Sprintf("%s『%s』 权限配置", baseTitle, model.Name)
	// e.Data["ColumnPowerSaveURL"] = fmt.Sprintf("/%s/ColumnPowerSave", BaseURL)
	// e.Data["URL_GetColumnPowerTreeViewJSON"] = fmt.Sprintf("/%s/GetColumnPowerTreeViewJSON", BaseURL)
	// e.Data["JSONTreeViewListSelectedDataUrl"] = fmt.Sprintf("/%s/GetTreeViewJSONSelectedData", BaseURL)

	// e.SetMortStype()
	// e.SetMortScript()
	// e.AppendCustomScripts([]string{
	// 	//Bootstrap table
	// 	"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
	// 	"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
	// 	// TreeView
	// 	"/static/js/hplus/plugins/treeview/bootstrap-treeview.min.js",
	// 	//layer
	// 	"/static/js/hplus/plugins/layer/layer.min.js",
	// })
	// e.AppendCustomStyles([]string{
	// 	//Bootstrap table
	// 	"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
	// 	// TreeView
	// 	"/static/css/hplus/treeview/bootstrap-treeview.min.css",
	// })

	// e.Layout = "admin/layout/layout.html"
	// e.LayoutSections = make(map[string]string)
	// e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	// e.TplName = BaseURL + "/columnpower.html"
}

// GetColumnPowerTreeViewJSON Default Json
func (e *UsersGroupController) GetColumnPowerTreeViewJSON() {
	// var list []*models.TreeView
	// var result models.Result

	// //取出用户组的所有权限
	// columnPowerList := admin.RoleS.GetPermissionsForUser("usergroup_" + e.GetString("id"))

	// // fmt.Println(columnPowerList)

	// list, err := new(models.ColumnPower).GetTreeViewBootstrap()
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// for _, v := range list {
	// 	for _, cv := range columnPowerList {
	// 		cid, _ := strconv.ParseInt(cv[1], 10, 64)
	// 		if v.ID == cid {
	// 			v.State = &models.TreeViewState{Checked: true, Expanded: true}
	// 		}
	// 		if v.Nodes != nil {
	// 			e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
	// 		}
	// 	}

	// }
	// // list[0].State = &models.TreeViewState{Checked: true, Expanded: true}

	// e.Data["json"] = list
	// e.ServeJSON()
}

// getRecursive 递归获取下一级
// func (e *UsersGroupController) getRecursiveColumnPowerTreeView(columnPowerList [][]string, list []*models.TreeView) {

// 	// for _, v := range list {
// 	// 	for _, cv := range columnPowerList {
// 	// 		cid, _ := strconv.ParseInt(cv[1], 10, 64)
// 	// 		if v.ID == cid {
// 	// 			v.State = &models.TreeViewState{Checked: true, Expanded: true}
// 	// 		}
// 	// 		if v.Nodes != nil {
// 	// 			e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
// 	// 		}
// 	// 	}
// 	// }
// }

// ColumnPowerSave 展示用户组-权限页面-保存
func (e *UsersGroupController) ColumnPowerSave() {
	// var result models.Result
	// idArray := strings.Split(e.GetString("ids"), ",")
	// id := e.GetString("id")
	// action := "GET"

	// //先删除所有的权限
	// admin.RoleS.RemoveFilteredPolicy(0, "usergroup_"+id)
	// // if !boo {
	// // 	result.Code = -1
	// // 	result.Msg = "删除权限失败"
	// // 	e.Data["json"] = result
	// // 	e.ServeJSON()
	// // 	return
	// // }

	// // return
	// // fmt.Println("admin.RoleS.GetPermissionsForUser", admin.RoleS.GetPermissionsForUser("usergroup_"+id))

	// for _, v := range idArray {
	// 	mid, _ := strconv.ParseInt(v, 10, 64)
	// 	m, _ := new(models.ColumnPower).GetOne(mid)
	// 	//保存用户组的权限
	// 	admin.RoleS.AddPolicy("usergroup_"+id, strconv.FormatInt(m.ID, 10), m.URL, action)
	// }

	// // admin.RoleS.LoadPolicy()
	// result.Code = 0
	// e.Data["json"] = result
	// e.ServeJSON()
}
