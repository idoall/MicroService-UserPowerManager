package models

import (
	"github.com/astaxie/beego/orm"
)

type CasbinRule struct {
	ID    int    `orm:"column(id);auto"`
	PType string `orm:"column(p_type);size(255);null" description:"g表示用户组，p表示权限"`
	V0    string `orm:"column(v0);size(255);null"`
	V1    string `orm:"column(v1);size(255);null"`
	V2    string `orm:"column(v2);size(255);null"`
	V3    string `orm:"column(v3);size(255);null"`
	V4    string `orm:"column(v4);size(45);null"`
	V5    string `orm:"column(v5);size(45);null"`
}

func (t *CasbinRule) TableName() string {
	return "casbin_rule"
}

func init() {
	orm.RegisterModel(new(CasbinRule))
}
