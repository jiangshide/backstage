package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int64
	Name       string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Motto      string
	Sex        int
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Admin) TableName() string {
	return TableName("uc_admin")
}

func (this *Admin) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Admin) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Admin) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Admin) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Admin) List(pageSize, offSet int) (list []*Admin, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}