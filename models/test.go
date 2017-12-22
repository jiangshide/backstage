package models

import "github.com/astaxie/beego/orm"

type Environment struct {
	Id         int64
	Name       string
	Jdk        string
	Git        string
	Gradle     string
	Adb        string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Environment) TableName() string {
	return TableName("test_environment")
}

func (this *Environment) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Environment) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Environment) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Environment) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Environment) List(pageSize, offSet int) (list []*Environment, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Project struct {
	Id         int64
	ParentId   int64
	Name       string
	Icon       string
	Address    string
	Account    string
	Psw        string
	Branch     string
	Tag        string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Project) TableName() string {
	return TableName("test_project")
}

func (this *Project) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Project) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Project) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Project) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Project) List(pageSize, offSet int) (list []*Project, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Test struct {
	Id         int64
	Name       string
	Type       int64
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Test) TableName() string {
	return TableName("test_test")
}

func (this *Test) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Test) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Test) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Test) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Test) List(pageSize, offSet int) (list []*Test, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
