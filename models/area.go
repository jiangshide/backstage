package models

import "github.com/astaxie/beego/orm"

type Continent struct {
	Id         int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Continent) TableName() string {
	return TableName("area_continent")
}

func (this *Continent) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Continent) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Continent) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Continent) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Continent) List(pageSize, offSet int) (list []*Continent, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type State struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *State) TableName() string {
	return TableName("area_state")
}

func (this *State) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *State) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *State) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *State) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *State) List(pageSize, offSet int) (list []*State, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Province struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Province) TableName() string {
	return TableName("area_province")
}

func (this *Province) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Province) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Province) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Province) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Province) List(pageSize, offSet int) (list []*Province, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type City struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *City) TableName() string {
	return TableName("area_city")
}

func (this *City) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *City) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *City) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *City) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *City) List(pageSize, offSet int) (list []*City, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Region struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Region) TableName() string {
	return TableName("area_region")
}

func (this *Region) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Region) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Region) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Region) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Region) List(pageSize, offSet int) (list []*Region, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type County struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *County) TableName() string {
	return TableName("area_county")
}

func (this *County) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *County) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *County) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *County) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *County) List(pageSize, offSet int) (list []*County, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Town struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Town) TableName() string {
	return TableName("area_town")
}

func (this *Town) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Town) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Town) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Town) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Town) List(pageSize, offSet int) (list []*Town, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Country struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Country) TableName() string {
	return TableName("area_country")
}

func (this *Country) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Country) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Country) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Country) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Country) List(pageSize, offSet int) (list []*Country, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Village struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Village) TableName() string {
	return TableName("area_village")
}

func (this *Village) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Village) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Village) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Village) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Village) List(pageSize, offSet int) (list []*Village, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Group struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Group) TableName() string {
	return TableName("area_group")
}

func (this *Group) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Group) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Group) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Group) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Group) List(pageSize, offSet int) (list []*Group, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Team struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Team) TableName() string {
	return TableName("area_team")
}

func (this *Team) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Team) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Team) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Team) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Team) List(pageSize, offSet int) (list []*Team, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
