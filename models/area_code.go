package models

import "github.com/astaxie/beego/orm"

type AreaCode struct {
	Id             int64
	NameCn         string
	NameEn         string
	AreaCode       string
	TimeDifference string
	Icon           string
	CreateId       int64
	UpdateId       int64
	CreateTime     int64
	UpdateTime     int64
}

func (this *AreaCode) TableName() string {
	return TableName("nation")
}

func (this *AreaCode) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *AreaCode) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *AreaCode) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *AreaCode) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *AreaCode) List(pageSize, offSet int) (list []*Nation, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
