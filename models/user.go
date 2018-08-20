package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int64
	Name  string
	Password   string
	RoleIds    string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *User) TableName() string {
	return TableName("uc_user")
}

func Add(this *User) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) GetByName(loginName string) (*User, error) {
	a := new(User)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (this *User) GetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(this.TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (this *User) GetById(id int64) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (this *User) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this);
	}
	return orm.NewOrm().Read(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *User) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

type Profile struct {
	Id              int64
	UserId          int64
	Icon            string
	Level           int
	Score           int8
	Name            string
	Phone           string
	Motto           string
	Sex             int8
	Email           string
	Weixin          string
	Qq              string
	Weibo           string
	IdCard          string
	IdCardPicFront  string
	IdCardPicBehind string
	CreateId        int64
	UpdateId        int64
	CreateTime      int64
	UpdateTime      int64
}

func (this *Profile) TableName() string {
	return TableName("uc_user_profile")
}

func (this *Profile) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Profile) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Profile) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Profile) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Profile) List(pageSize, offSet int) (list []*Profile, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Device struct {
	Id         int64
	UserId     int64
	Name       string
	Mac        string
	Latitude   string
	Longitude  string
	Mode       string
	Arch       string
	SdkVersion string
	AppVersion string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Device) TableName() string {
	return TableName("uc_user_device")
}

func (this *Device) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Device) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Device) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Device) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Device) List(pageSize, offSet int) (list []*Device, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Brand struct {
	Id         int64
	UserId     int64
	Icon       string
	Name       string
	Code       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Brand) TableName() string {
	return TableName("uc_user_brand")
}

func (this *Brand) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Brand) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Brand) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Brand) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Brand) List(pageSize, offSet int) (list []*Brand, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Address struct {
	Id         int64
	UserId     int64
	Name       string
	Code       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Address) TableName() string {
	return TableName("uc_user_address")
}

func (this *Address) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Address) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Address) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Address) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Address) List(pageSize, offSet int) (list []*Address, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
