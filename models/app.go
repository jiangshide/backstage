package models

import (
	"github.com/astaxie/beego/orm"
)

type Channel struct {
	Id         int64
	Name       string
	FriendId   string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Channel) TableName() string {
	return TableName("app_channel")
}

func (this *Channel) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Channel) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Channel) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Channel) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Channel) List(pageSize, offSet int) (list []*Channel, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type AppName struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *AppName) TableName() string {
	return TableName("app_name")
}

func (this *AppName) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *AppName) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *AppName) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *AppName) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *AppName) List(pageSize, offSet int) (list []*AppName, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Pkgs struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Pkgs) TableName() string {
	return TableName("app_package")
}

func (this *Pkgs) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Pkgs) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Pkgs) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Pkgs) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Pkgs) List(pageSize, offSet int) (list []*Pkgs, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Version struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Version) TableName() string {
	return TableName("app_version")
}

func (this *Version) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Version) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Version) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Version) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Version) List(pageSize, offSet int) (list []*Version, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Code struct {
	Id         int64
	Code       int
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Code) TableName() string {
	return TableName("app_code")
}

func (this *Code) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Code) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Code) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Code) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Code) List(pageSize, offSet int) (list []*Code, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Env struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Env) TableName() string {
	return TableName("app_env")
}

func (this *Env) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Env) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Env) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Env) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Env) List(pageSize, offSet int) (list []*Env, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Build struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Build) TableName() string {
	return TableName("app_build")
}

func (this *Build) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Build) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Build) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Build) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Build) List(pageSize, offSet int) (list []*Build, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Type struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Type) TableName() string {
	return TableName("app_type")
}

func (this *Type) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Type) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Type) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Type) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Type) List(pageSize, offSet int) (list []*Type, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type App struct {
	Id            int64
	ProjectId     int64
	TestId        int64
	Icon          string
	TypeId        int64
	ApplicationId int64
	PkgId         int64
	VersionId     int64
	CodeId        int64
	EnvId         int64
	BuildId       int64
	ChannelId     int64
	Descript      string
	Status        int
	Times         int
	Url           string
	Downs         int
	CreateId      int64
	UpdateId      int64
	CreateTime    int64
	UpdateTime    int64
	Views         int
}

func (this *App) TableName() string {
	return TableName("app")
}

func (this *App) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *App) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *App) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *App) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *App) List(pageSize, offSet int) (list []*App, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
