package models

import (
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	Icon       string `json:"icon"`
	Descript   string `json:"descript"`
	Clicks     int    `json:"clicks"`
	CreateId   int64  `json:"createId"`
	UpdateId   int64  `json:"UpdateId"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Views      int    `json:"views"`
}

func (this *Banner) TableName() string {
	return TableName("web_banner")
}

func (this *Banner) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Banner) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Banner) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Banner) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Banner) List(pageSize, offSet int) (list []*Banner, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
