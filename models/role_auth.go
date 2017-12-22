package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"bytes"
	"strconv"
)

type RoleAuth struct {
	AuthId int64 `orm:"pk"`
	RoleId int64
}

func (this *RoleAuth) TableName() string {
	return TableName("uc_role_auth")
}

func (this *RoleAuth) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *RoleAuth) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *RoleAuth) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *RoleAuth) Query() error {
	if this.AuthId == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *RoleAuth) List() (list []*RoleAuth, err error) {
	_, err = orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).All(&list)
	return
}

func RoleAuthGetByIds(roleIds string) (string, error) {
	list := make([]*RoleAuth, 0)
	ids := strings.Split(roleIds, ",")
	_, err := orm.NewOrm().QueryTable(TableName("uc_role_auth")).Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.FormatInt(v.AuthId, 11))
			b.WriteString(",")
		}
	}
	return strings.TrimRight(b.String(), ","), nil
}

func (this *RoleAuth) BatchAdd(ras []*RoleAuth) (n int, err error) {
	query := orm.NewOrm().QueryTable(this.TableName())
	insert, _ := query.PrepareInsert()
	defer insert.Close()
	for _, ra := range ras {
		if _, err = insert.Insert(ra); err == nil {
			n += 1
		}
	}
	return
}
