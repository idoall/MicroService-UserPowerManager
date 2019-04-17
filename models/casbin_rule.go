package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CasbinRule struct {
	Id    int    `orm:"column(id);auto"`
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

// AddCasbinRule insert a new CasbinRule into database and returns
// last inserted Id on success.
func AddCasbinRule(m *CasbinRule) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCasbinRuleById retrieves CasbinRule by Id. Returns error if
// Id doesn't exist
func GetCasbinRuleById(id int) (v *CasbinRule, err error) {
	o := orm.NewOrm()
	v = &CasbinRule{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCasbinRule retrieves all CasbinRule matches certain condition. Returns empty list if
// no records exist
func GetAllCasbinRule(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CasbinRule))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CasbinRule
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCasbinRule updates CasbinRule by Id and returns error if
// the record to be updated doesn't exist
func UpdateCasbinRuleById(m *CasbinRule) (err error) {
	o := orm.NewOrm()
	v := CasbinRule{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCasbinRule deletes CasbinRule by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCasbinRule(id int) (err error) {
	o := orm.NewOrm()
	v := CasbinRule{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CasbinRule{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
