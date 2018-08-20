package models

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RoleAuth struct {
	AuthId int64 `orm:"pk"`
	RoleId int64
}

func (this *RoleAuth) TableName() string {
	return TableName("uc_role_auth")
}

func (this *RoleAuth)Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *RoleAuth)GetById(id int64) ([]*RoleAuth, error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(this.TableName())
	_, err := query.Filter("role_id", id).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (this *RoleAuth)Delete(id int64) (int64, error) {
	query := orm.NewOrm().QueryTable(this.TableName())
	return query.Filter("role_id", id).Delete()
}

//获取多个
func RoleAuthGetByIds(RoleIds string) (Authids string, err error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_role_auth"))
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.FormatInt(v.AuthId,10))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

func RoleAuthMultiAdd(this []*RoleAuth) (n int, err error) {
	query := orm.NewOrm().QueryTable(TableName("uc_role_auth"))
	i, _ := query.PrepareInsert()
	for _, ra := range this {
		_, err := i.Insert(ra)
		if err == nil {
			n = n + 1
		}
	}
	i.Close() // 别忘记关闭 statement
	return n, err
}