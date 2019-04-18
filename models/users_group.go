package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

type UsersGroup struct {
	Id             int       `orm:"column(id);auto"`
	Name           string    `orm:"column(name);size(200);null" description:"用户组名称"`
	ParentId       int       `orm:"column(parent_id)" description:"所属组Id"`
	Sorts          int       `orm:"column(sorts)" description:"排序"`
	Note           string    `orm:"column(note);size(2000);null" description:"备注"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime)"`
	LastUpdateTime time.Time `orm:"column(last_update_time);type(datetime);null" description:"最后更新时间"`
}

func (t *UsersGroup) TableName() string {
	return "users_group"
}

func init() {
	orm.RegisterModel(new(UsersGroup))
}

// Add 添加
func (e *UsersGroup) Add(k *UsersGroup) (int64, error) {
	o := orm.NewOrm()
	k.CreateTime = time.Now()
	k.LastUpdateTime = k.CreateTime
	return o.Insert(k)
}

// GetOne 获取一条记录
func (e *UsersGroup) GetOne(id int64) (*UsersGroup, error) {
	return e.QueryOne(orm.NewCondition().And("id", id), "-id")
}

// GetChildIdArray 获取创建的用户ID列表
func (e *UsersGroup) GetChildIdArray(id int64) ([]int64, error) {
	cond := orm.NewCondition().And("parent_id", id)
	list, _, err := e.GetAll(cond, 1000, 1, "-id")
	if err != nil && !commonutils.StringContains(err.Error(), "no row found") {
		return nil, err
	}
	result := []int64{}
	for _, v := range list {
		result = append(result, int64(v.Id))
	}
	return result, nil
}

// QueryOne 获取一条记录
func (e *UsersGroup) QueryOne(cond *orm.Condition, order ...string) (*UsersGroup, error) {
	o := orm.NewOrm()
	result := UsersGroup{}
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
func (e *UsersGroup) GetAll(cond *orm.Condition, pageSize, currentPageIndex int, order ...string) ([]*UsersGroup, int64, error) {
	o := orm.NewOrm()
	var resultlist []*UsersGroup
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
func (e *UsersGroup) RawSQL(sql string, param ...interface{}) (sql.Result, error) {
	o := orm.NewOrm()
	return o.Raw(sql, param).Exec()
}

// Update 修改
func (e *UsersGroup) Update(k *UsersGroup) (int64, error) {
	o := orm.NewOrm()
	k.LastUpdateTime = time.Now()
	return o.Update(k)
}

// Delete  删除一条记录
func (e *UsersGroup) Delete(id int64) (int64, error) {

	return e.Delete(id)
}

// BatchDelete  批量删除多条记录
func (e *UsersGroup) BatchDelete(id []string) (sql.Result, error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", e.TableName(), strings.Join(id, ","))
	return e.RawSQL(sql)
}
