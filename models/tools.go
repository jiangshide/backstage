package models

import "github.com/astaxie/beego/orm"

type Qrcode struct {
	Id         int64
	Name       string
	Content    string
	Url        string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Qrcode) TableName() string {
	return TableName("tools_qrcode")
}

func (this *Qrcode) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Qrcode) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Qrcode) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Qrcode) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Qrcode) List(pageSize, offSet int) (list []*Qrcode, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type FormatType struct {
	Id         int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *FormatType) TableName() string {
	return TableName("tools_format_type")
}

func (this *FormatType) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *FormatType) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *FormatType) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *FormatType) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *FormatType) List(pageSize, offSet int) (list []*FormatType, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Format struct {
	Id         int64
	ParentId   int64
	Name       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Format) TableName() string {
	return TableName("tools_format")
}

func (this *Format) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Format) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Format) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Format) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Format) List(pageSize, offSet int) (list []*Format, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Compress struct {
	Id         int64
	Name       string
	Url        string
	Type       int64
	Format     string
	Descript   string
	Size       int64
	ReSize     int64
	Compress   int
	Downs      int
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Compress) TableName() string {
	return TableName("tools_compress")
}

func (this *Compress) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Compress) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Compress) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Compress) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Compress) List(pageSize, offSet int) (list []*Compress, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
