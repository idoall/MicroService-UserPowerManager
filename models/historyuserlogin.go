package models

import (
	"database/sql"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/astaxie/beego/orm"
)

// HistoryUserLogin 历史记录-用户登录
type HistoryUserLogin struct {
	ID             int64     `orm:"column(id);unique" form:"-"` // 编号id
	User           *Users    `orm:"column(user_id);rel(one)"`
	DeviceDetector string    `orm:"column(device_detector)"` //设备检测器
	GeoRemoteAddr  string    `orm:"column(geo_remote_addr)"`
	GeoCountry     string    `orm:"column(geo_country)"`
	GeoCity        string    `orm:"column(geo_city)"`
	CreateTime     time.Time `orm:"column(create_time);auto_now;type(datetime)"` //加入时间
}

// TableName 表的名称
func (e *HistoryUserLogin) TableName() string {
	return "history_user_login"
}

func init() {
	orm.RegisterModel(new(HistoryUserLogin))
}

// Add 添加
func (e *HistoryUserLogin) Add(k *HistoryUserLogin) (int64, error) {
	o := orm.NewOrm()
	k.CreateTime = time.Now()
	return o.Insert(k)
}

// GetOne 获取一条记录
func (e *HistoryUserLogin) GetOne(id int64) (*HistoryUserLogin, error) {
	return e.QueryOne(orm.NewCondition().And("id", id), "-id")
}

// GetChildIdArray 获取创建的用户ID列表
func (e *HistoryUserLogin) GetChildIdArray(id int64) ([]int64, error) {
	cond := orm.NewCondition().And("parent_id", id)
	list, _, err := e.GetAll(cond, 1000, 1, "-id")
	if err != nil && !commonutils.StringContains(err.Error(), "no row found") {
		return nil, err
	}
	result := []int64{}
	for _, v := range list {
		result = append(result, v.ID)
	}
	return result, nil
}

// QueryOne 获取一条记录
func (e *HistoryUserLogin) QueryOne(cond *orm.Condition, order ...string) (*HistoryUserLogin, error) {
	o := orm.NewOrm()
	result := HistoryUserLogin{}
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
func (e *HistoryUserLogin) GetAll(cond *orm.Condition, pageSize, currentPageIndex int, order ...string) ([]*HistoryUserLogin, int64, error) {
	o := orm.NewOrm()
	var resultlist []*HistoryUserLogin
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
func (e *HistoryUserLogin) RawSQL(sql string, param ...interface{}) (sql.Result, error) {
	o := orm.NewOrm()
	return o.Raw(sql, param).Exec()
}
