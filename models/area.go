package models

import "github.com/astaxie/beego/orm"

type Area struct {
	Id             int64
	ParentId       int64  //所属ID:0~洲,1~国家,2~省,3~市,4~地区,5~县,6~镇,7~乡,8~村,9~组,10~队
	Name           string //名称
	NameAb         string //名称简写
	Icon           string //Logo
	NameCn         string //代码名称~中文:中国
	NameEn         string //代码名称~外文:china
	ZoneCode       int    //区号:86
	ZipCode        int    //邮编:999001
	AreaCode       string //区域代码～北京:1000
	TimeDifference string //时差:0
	CreateId       int64
	UpdateId       int64
	CreateTime     int64
	UpdateTime     int64
}

func (this *Area) TableName() string {
	return TableName("area")
}

func (this *Area) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Area) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Area) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Area) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Area) List(pageSize, offSet int) (list []*Area, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
