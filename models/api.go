package models

import "github.com/astaxie/beego/orm"

type Key struct {
	Id         int64
	Name       string
	Key        string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Key) TableName() string {
	return TableName("api_key")
}

func (this *Key) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Key) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Key) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Key) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Key) List(pageSize, offSet int) (list []*Key, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
