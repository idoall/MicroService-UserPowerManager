package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// swagger:model
type Users struct {
	// 用户Id
	// required: false
	// Read Only
	Id int `orm:"column(id);auto" description:"编号Id"`
	// 用户名称
	// required: true
	// min length: 3
	UserName  string `orm:"column(user_name);size(200);null" description:"用户名称"`
	RealyName string `orm:"column(realy_name);size(200);null" description:"真实姓名"`
	// 用户的密码
	// required: true
	// min length: 6
	Password string `orm:"column(password);size(100)" description:"密码"`
	// required: false
	// Read Only
	AuthKey string `orm:"column(auth_key);size(45);null" description:"authkey"`
	// 用户的Email地址
	// required: true
	// example: user@provider.net
	Email string `orm:"column(email);size(100);null" description:"email"`
	// swagger:ignore
	IsDel          bool      `orm:"column(is_del);null" description:"是否删除：1、true 0、false"`
	Note           string    `orm:"column(note);null" description:"加入时间"`
	ParentId       int       `orm:"column(parent_id)"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime)"`
	LastUpdateTime time.Time `orm:"column(last_update_time);type(datetime);null" description:"最后更新时间"`
}

// TableName 表的名称
func (e *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// Add 添加
func (e *Users) Add(k *Users) (int64, error) {
	o := orm.NewOrm()
	k.CreateTime = time.Now()
	k.LastUpdateTime = k.CreateTime
	return o.Insert(k)
}

// GetOne 获取一条记录
func (e *Users) GetOne(id int64) (*Users, error) {
	return e.QueryOne(orm.NewCondition().And("id", id), "-id")
}

// GetChildIdArray 获取创建的用户ID列表
func (e *Users) GetChildIdArray(id int64) ([]int64, error) {
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
func (e *Users) QueryOne(cond *orm.Condition, order ...string) (*Users, error) {
	o := orm.NewOrm()
	result := Users{}
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
func (e *Users) GetAll(cond *orm.Condition, pageSize, currentPageIndex int, order ...string) ([]*Users, int64, error) {
	o := orm.NewOrm()
	var resultlist []*Users
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
func (e *Users) RawSQL(sql string, param ...interface{}) (sql.Result, error) {
	o := orm.NewOrm()
	return o.Raw(sql, param).Exec()
}

// Update 修改
func (e *Users) Update(k *Users) (int64, error) {
	o := orm.NewOrm()
	k.LastUpdateTime = time.Now()
	return o.Update(k)
}

// Delete  删除一条记录
func (e *Users) Delete(id int64) (int64, error) {

	model, _ := e.GetOne(id)
	model.IsDel = true
	model.LastUpdateTime = time.Now()
	return e.Update(model)
}

// BatchDelete  批量删除多条记录
func (e *Users) BatchDelete(id []string) (sql.Result, error) {
	sql := fmt.Sprintf("UPDATE %s SET is_del=1,last_update_time=now() WHERE id IN (%s)", e.TableName(), strings.Join(id, ","))
	return e.RawSQL(sql)
}
