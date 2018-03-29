package models

import "github.com/astaxie/beego/orm"

type UserProfile struct {
	Id           int64
	UserId       int64
	PswIntensity int
	Icon         string
	Level        int
	Score        int
	Name         string
	Phone        string
	Motto        string
	Sex          int
	Email        string
	Weixin       string
	Qq           string
	Weibo        string
	Address      string
	Code         int
	CreateId     int64
	UpdateId     int64
	CreateTime   int64
	UpdateTime   int64
}

func (this *UserProfile) TableName() string {
	return TableName("user_profile")
}

func (this *UserProfile) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *UserProfile) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *UserProfile) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *UserProfile) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *UserProfile) List(pageSize, offSet int) (list []*UserProfile, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
