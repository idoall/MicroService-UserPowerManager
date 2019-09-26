package models

import (
	"database/sql"
	"time"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"

	"github.com/go-xorm/builder"

	"github.com/idoall/MicroService-UserPowerManager/utils/orm"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

const (
	selectSQL = "id, user_name, realy_name, password, auth_key, email, is_del, note, parent_id, create_time, last_update_time"
)

// Users ORM
// 用户表使用了 KingShared 的分表方案
type Users struct {
	Id             int64     `xorm:"pk comment('编号Id') BIGINT(20)"`
	UserName       string    `xorm:"comment('用户名称') VARCHAR(200)"`
	RealyName      string    `xorm:"comment('真实姓名') VARCHAR(200)"`
	Password       string    `xorm:"not null comment('密码') VARCHAR(100)"`
	AuthKey        string    `xorm:"comment('authkey') VARCHAR(45)"`
	Email          string    `xorm:"comment('email') VARCHAR(100)"`
	IsDel          int       `xorm:"default 0 comment('是否删除：1、true 0、false') TINYINT(1)"`
	Note           string    `xorm:"comment('加入时间') TEXT"`
	ParentId       int64     `xorm:"not null BIGINT(20)"`
	CreateTime     time.Time `xorm:"not null DATETIME"`
	LastUpdateTime time.Time `xorm:"comment('最后更新时间') DATETIME"`
}

// TableName 表的名称
func (e *Users) TableName() string {
	return "users"
}

// Add 添加
func (e *Users) Add(k *Users) (int64, error) {
	var counts int64
	var err error

	o := orm.XROM

	// 短线方案取数据库最大记录数，量大性能不好，可以换方案
	if counts, err = o.Count(&Users{}); err != nil {
		inner.Mlogger.Fatal(err)
	}

	k.Id = counts + 1
	k.CreateTime = time.Now()
	k.LastUpdateTime = time.Now()
	if _, err = o.Insert(k); err != nil {
		return 0, err
	}

	return k.Id, nil

}

// GetOne 获取一条记录
func (e *Users) GetOne(id int64) (Users, error) {
	return e.QueryOne(builder.Eq{"id": id}, "id desc")
}

// GetChildIDArray 获取创建的用户ID列表
func (e *Users) GetChildIDArray(id int64) ([]int64, error) {

	list, _, err := e.GetAll(builder.Eq{"parent_id": id}, "id desc", 1000, 1)
	if err != nil && !commonutils.StringContains(err.Error(), "no row found") {
		return nil, err
	}
	result := []int64{}
	for _, v := range list {
		result = append(result, v.Id)
	}
	return result, nil
}

// QueryOne 获取一条记录
// u := &models.Users{}
// u.QueryOne(Eq{"id": 1}, "id desc")
func (e *Users) QueryOne(whereCond builder.Cond, orderBy string) (Users, error) {

	if resultlist, _, err := e.GetAll(whereCond, orderBy, 1, 1); err != nil {
		return Users{}, err
	} else if len(resultlist) == 0 {
		return Users{}, nil
	} else {
		return resultlist[0], err
	}
}

// GetAll 获取
func (e *Users) GetAll(whereCond builder.Cond, orderBy string, pageSize, currentPageIndex int) ([]Users, int64, error) {

	var err error
	var count int64
	var whereSQL string
	o := orm.XROM

	resultlist := []Users{}

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
func (e *Users) RawSQL(sql ...interface{}) (sql.Result, error) {
	o := orm.XROM
	return o.Exec(sql)
}

// Update 修改
func (e *Users) Update(k Users) (int64, error) {
	o := orm.XROM

	// 默认只更新非空和非0的字段
	updateID := k.Id
	k.Id = 0

	// 设置更新时间
	k.LastUpdateTime = time.Now()

	return o.ID(updateID).Update(k)
}

// Delete  删除一条记录
func (e *Users) Delete(id int64) (int64, error) {

	model, _ := e.GetOne(id)
	model.IsDel = 1
	return e.Update(model)
}

// BatchDelete  批量删除多条记录
func (e *Users) BatchDelete(IDArray []int64) (int64, error) {
	o := orm.XROM
	return o.In("id", IDArray).Update(&Users{IsDel: 1})
}
