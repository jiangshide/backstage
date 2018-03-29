package models

import (
	"github.com/astaxie/beego/orm"
)

type PermissionUser struct {
	Id         int64
	Name       string
	Password   string
	RoleId     string
	Salt       string
	Status     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *PermissionUser) TableName() string {
	return TableName("permission_user")
}

func (this *PermissionUser) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *PermissionUser) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *PermissionUser) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *PermissionUser) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *PermissionUser) List(pageSize, offSet int) (list []*PermissionUser, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
