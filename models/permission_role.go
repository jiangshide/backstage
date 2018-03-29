package models

import (
	"github.com/astaxie/beego/orm"
)

type PermissionRole struct {
	Id         int64
	Name       string
	Detail     string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *PermissionRole) TableName() string {
	return TableName("permission_role")
}

func (this *PermissionRole) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *PermissionRole) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *PermissionRole) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *PermissionRole) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *PermissionRole) List(pageSize, offSet int) (list []*PermissionRole, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
