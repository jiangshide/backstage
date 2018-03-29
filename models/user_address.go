package models

import "github.com/astaxie/beego/orm"

type UserAddress struct {
	Id         int64
	UserId     int64
	Name       string
	Code       int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *UserAddress) TableName() string {
	return TableName("user_address")
}

func (this *UserAddress) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *UserAddress) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *UserAddress) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *UserAddress) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *UserAddress) List(pageSize, offSet int) (list []*UserAddress, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
