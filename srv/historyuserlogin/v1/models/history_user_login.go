package models

import (
	"time"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/orm"
)

// HistoryUserLogin struct
type HistoryUserLogin struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	UserId         int64     `xorm:"not null comment('用户ID')"`
	GeoRemoteAddr  string    `xorm:"not null comment('用户登录IP') VARCHAR(50)"`
	GeoCountry     string    `xorm:"not null comment('IP所在国家') VARCHAR(100)"`
	GeoCity        string    `xorm:"not null comment('IP所在城市') VARCHAR(100)"`
	DeviceDetector string    `xorm:"not null comment('设备检测器') VARCHAR(1000)"`
	CreateTime     time.Time `xorm:"not null comment('创建时间') DATETIME"`
}

// TableName 表的名称
func (e *HistoryUserLogin) TableName() string {
	return "history_user_login"
}

// Add 添加
func (e *HistoryUserLogin) Add(m HistoryUserLogin) (int64, error) {
	var err error

	o := orm.XROM

	m.CreateTime = time.Now()
	if _, err = o.Insert(m); err != nil {
		if utils.RunMode == "dev" {
			inner.Mlogger.Errorf("[HistoryUserLogin][Add] %s", err.Error())
		}
		return 0, err
	}

	return 1, nil

}
