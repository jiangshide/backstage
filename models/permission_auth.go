package models

import (
	"github.com/astaxie/beego/orm"
)

type PermissionAuth struct {
	Id         int64
	Pid        int
	Name       string
	Action     string
	Icon       string
	IsShow     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *PermissionAuth) TableName() string {
	return TableName("permission_auth")
}

func (this *PermissionAuth) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *PermissionAuth) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *PermissionAuth) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *PermissionAuth) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *PermissionAuth) QueryAll() (list []*PermissionAuth, err error) {
	_, err = orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).All(&list)
	return
}

func (this *PermissionAuth) List(pageSize, offSet int) (list []*PermissionAuth, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
