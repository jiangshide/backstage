package models

import "github.com/astaxie/beego/orm"

type UserLocation struct {
	Id         int64
	UserId     int64
	Ip         string
	Mac        string
	Latitude   string
	Longitude  string
	Device     string
	Mode       string
	Arch       string
	SdkVersion string
	AppVersion string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *UserLocation) TableName() string {
	return TableName("user_location")
}

func (this *UserLocation) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *UserLocation) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *UserLocation) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *UserLocation) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *UserLocation) List(pageSize, offSet int) (list []*UserLocation, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
