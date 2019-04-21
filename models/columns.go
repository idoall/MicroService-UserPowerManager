package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

type Columns struct {
	ID             int64      `orm:"column(id);auto"`
	Name           string     `orm:"column(name);size(200)" description:"栏目名称"`
	URL            string     `orm:"column(URL);size(500);null" description:"URL"`
	ParentID       int64      `orm:"column(parent_id)" description:"所属上级Id"`
	Sorts          int64      `orm:"column(sorts)" description:"排序"`
	IsShowNav      bool       `orm:"column(is_show_nav)" description:"是否显示在导航"`
	CssIcon        string     `orm:"column(css_icon);size(50);null" description:"css图标样式"`
	Nodes          []*Columns `orm:"-" json:"nodes"` //下属节点集合
	CreateTime     time.Time  `orm:"column(create_time);type(datetime)"`
	LastUpdateTime time.Time  `orm:"column(last_update_time);type(datetime);null" description:"最后更新时间"`
}

func (t *Columns) TableName() string {
	return "columns"
}

func init() {
	orm.RegisterModel(new(Columns))
}

// Add 添加
func (e *Columns) Add(k *Columns) (int64, error) {
	o := orm.NewOrm()
	k.CreateTime = time.Now()
	k.LastUpdateTime = k.CreateTime
	return o.Insert(k)
}

// GetOne 获取一条记录
func (e *Columns) GetOne(id int64) (*Columns, error) {
	return e.QueryOne(orm.NewCondition().And("id", id), "-id")
}

// GetChildIdArray 获取创建的用户ID列表
func (e *Columns) GetChildIdArray(id int64) ([]int64, error) {
	cond := orm.NewCondition().And("parent_id", id)
	list, _, err := e.GetAll(cond, 1000, 1, "-id")
	if err != nil && !commonutils.StringContains(err.Error(), "no row found") {
		return nil, err
	}
	result := []int64{}
	for _, v := range list {
		result = append(result, int64(v.ID))
	}
	return result, nil
}

// QueryOne 获取一条记录
func (e *Columns) QueryOne(cond *orm.Condition, order ...string) (*Columns, error) {
	o := orm.NewOrm()
	result := Columns{}
	qs := o.QueryTable(e.TableName())

	qs = qs.SetCond(cond)

	qs = qs.RelatedSel()

	//Order
	if len(order) != 0 {
		qs = qs.OrderBy(order...)
	}

	err := qs.One(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// GetAll 获取
func (e *Columns) GetAll(cond *orm.Condition, pageSize, currentPageIndex int, order ...string) ([]*Columns, int64, error) {
	o := orm.NewOrm()
	var resultlist []*Columns
	var count int64

	qs := o.QueryTable(e.TableName())

	qs = qs.SetCond(cond)

	qs = qs.RelatedSel()
	//Order
	if len(order) != 0 {
		qs = qs.OrderBy(order...)
	}

	_, err := qs.Limit(pageSize, (currentPageIndex-1)*pageSize).All(&resultlist)
	if err != nil {
		return resultlist, count, err
	}
	count, err = qs.Count()
	if err != nil {
		return resultlist, count, err
	}

	return resultlist, count, nil
}

// RawSQL 执行sql
func (e *Columns) RawSQL(sql string, param ...interface{}) (sql.Result, error) {
	o := orm.NewOrm()
	return o.Raw(sql, param).Exec()
}

// Update 修改
func (e *Columns) Update(k *Columns) (int64, error) {
	o := orm.NewOrm()
	k.LastUpdateTime = time.Now()
	return o.Update(k)
}

// Delete  删除一条记录
func (e *Columns) Delete(id int64) (int64, error) {

	return e.Delete(id)
}

// BatchDelete  批量删除多条记录
func (e *Columns) BatchDelete(id []string) (sql.Result, error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", e.TableName(), strings.Join(id, ","))
	return e.RawSQL(sql)
}

// GetTreeViewBootstrap GetTreeViewBootstrap
// func (e *Columns) GetTreeViewBootstrap() ([]*TreeView, error) {
// 	// 根据 sorts获取所有集合倒排序
// 	list, _, err := e.GetAll(orm.NewCondition(), 1000, 1, "-sorts")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var returnlist []*TreeView
// 	for _, v := range list {
// 		// 取出 Id 为0的
// 		if v.ParentID == 0 {
// 			t := &TreeView{
// 				ID:   v.ID,
// 				Text: v.Name,
// 				Href: v.URL,
// 				Tags: []string{fmt.Sprintf("%d", v.ID), v.URL, fmt.Sprintf("Sort:%d", v.Sort), fmt.Sprintf("IsShowNav:%t", v.IsShowNav)},
// 			}
// 			e.generateRecursiveTreeViewBootstrap(v, t, list)
// 			returnlist = append(returnlist, t)
// 		}
// 	}
// 	return returnlist, nil
// }

// func (e *Columns) generateRecursiveTreeViewBootstrap(m *ColumnPower, t *TreeView, list []*ColumnPower) {

// 	for _, v := range list {
// 		if v.ParentID == m.ID {
// 			tr := &TreeView{
// 				ID:   v.ID,
// 				Text: v.Name,
// 				Href: v.URL,
// 				Tags: []string{fmt.Sprintf("%d", v.ID), v.URL, fmt.Sprintf("Sort:%d", v.Sort), fmt.Sprintf("IsShowNav:%t", v.IsShowNav)},
// 			}
// 			t.Nodes = append(t.Nodes, tr)
// 			e.generateRecursiveTreeViewBootstrap(v, tr, list)
// 		}
// 	}
// }
