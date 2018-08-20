package models

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int64
	RoleName   string
	Detail     string
	Status     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Role) TableName() string {
	return TableName("uc_role")
}

func (this *Role)GetList(page, pageSize int, filters ...interface{}) ([]*Role, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Role, 0)
	query := orm.NewOrm().QueryTable(this.TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (this *Role)Add() (int64, error) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *Role)GetById(id int64) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (this *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Role) Del()(int64,error){
	return orm.NewOrm().Delete(this)
}