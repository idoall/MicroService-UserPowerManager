package models

import (
	"database/sql"
	"time"

	"github.com/go-xorm/builder"
	"github.com/idoall/MicroService-UserPowerManager/utils/orm"
)

const (
	selectSQL = "id, name, URL, parent_id, sorts, is_show_nav, css_icon, create_time, last_update_time"
)

// Columns struct
type Columns struct {
	Id             int64     `xorm:"not null pk autoincr unique INT(11)"`
	Name           string    `xorm:"not null comment('栏目名称') VARCHAR(200)"`
	Url            string    `xorm:"comment('URL') VARCHAR(500)"`
	ParentId       int64     `xorm:"not null comment('所属上级Id') INT(11)"`
	Sorts          int32     `xorm:"not null default 0 comment('排序') INT(11)"`
	IsShowNav      int32     `xorm:"not null default 0 comment('是否显示在导航') INT(1)"`
	CssIcon        string    `xorm:"comment('css图标样式') VARCHAR(50)"`
	CreateTime     time.Time `xorm:"not null comment('创建时间') DATETIME"`
	LastUpdateTime time.Time `xorm:"comment('最后更新时间') DATETIME"`
}

// TableName 表的名称
func (e *Columns) TableName() string {
	return "columns"
}

// Add 添加
func (e *Columns) Add(m *Columns) (int64, error) {
	o := orm.XROM
	m.CreateTime = time.Now()
	m.LastUpdateTime = m.CreateTime
	if _, err := o.Insert(m); err != nil {
		return 0, err
	}

	return 1, nil

}

// GetOne 获取一条记录
func (e *Columns) GetOne(id int64) (Columns, error) {
	return e.QueryOne(builder.Eq{"id": id}, "id desc")
}

// QueryOne 获取一条记录
func (e *Columns) QueryOne(whereCond builder.Cond, orderBy string) (Columns, error) {

	if resultlist, _, err := e.GetAll(whereCond, orderBy, 1, 1); err != nil {
		return Columns{}, err
	} else if len(resultlist) == 0 {
		return Columns{}, nil
	} else {
		return resultlist[0], err
	}
}

// GetAll 获取
func (e *Columns) GetAll(whereCond builder.Cond, orderBy string, pageSize, currentPageIndex int) ([]Columns, int64, error) {

	var err error
	var count int64
	var whereSQL string
	o := orm.XROM

	resultlist := []Columns{}

	// 拼接 SQL 查询语句
	if whereSQL, err = builder.Dialect(builder.MYSQL).Select(selectSQL).From(e.TableName()).Where(whereCond).OrderBy(orderBy).Limit(pageSize, (currentPageIndex-1)*pageSize).ToBoundSQL(); err != nil {
		return resultlist, count, err
	}

	if err = o.Sql(whereSQL).Find(&resultlist); err != nil {
		count, err = o.Sql(whereSQL).Count()
		return resultlist, count, err
	}

	return resultlist, count, err
}

// RawSQL 执行sql
func (e *Columns) RawSQL(sql ...interface{}) (sql.Result, error) {
	o := orm.XROM
	return o.Exec(sql)
}

// Update 修改
func (e *Columns) Update(m Columns) (int64, error) {
	o := orm.XROM

	// 默认只更新非空和非0的字段
	updateID := m.Id
	m.Id = 0

	// 设置更新时间
	m.LastUpdateTime = time.Now()

	return o.ID(updateID).Update(m)
}

// Delete  删除一条记录
func (e *Columns) Delete(id int64) (int64, error) {
	return e.BatchDelete([]int64{id})
}

// BatchDelete  批量删除多条记录
func (e *Columns) BatchDelete(IDArray []int64) (int64, error) {
	o := orm.XROM
	return o.In("id", IDArray).Delete(&Columns{})
}
