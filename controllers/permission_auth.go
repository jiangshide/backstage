package controllers

import (
	"zd112_backstage/models"
	"time"
)

type PermissionAuthContrller struct {
	BaseController
}

func (this *PermissionAuthContrller) List() {
	this.pageTitle("权限因子列表")
	this.display(this.getBgPermissionAction("auth/list"))
}

func (this *PermissionAuthContrller) Add() {
	this.pageTitle("增加权限因子")
	this.display(this.getBgPermissionAction("auth/add"))
}

func (this *PermissionAuthContrller) Edit() {
	auth := new(models.PermissionAuth)
	auth.Id = this.getId64(0)
	if err := auth.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, auth, true, this.getBgPermissionAction("auth/edit"))
}

func (this *PermissionAuthContrller) AjaxSave() {
	auth := new(models.PermissionAuth)
	auth.Id = this.getId64(0)
	auth.Name = this.getString("name", "名称不能为空!", defaultMinSize)
	auth.Action = this.getString("action", "Action不能为空!", defaultMinSize)
	auth.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	auth.IsShow = this.getInt("is_show", 0)
	var err error
	if auth.Id == 0 {
		auth.CreateId = this.userId
		auth.CreateTime = time.Now().Unix()
		_, err = auth.Add()
	} else {
		auth.UpdateId = this.userId
		auth.UpdateTime = time.Now().Unix()
		_, err = auth.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *PermissionAuthContrller) Table() {
	auth := new(models.PermissionAuth)
	result, count := auth.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *PermissionAuthContrller) AjaxDel() {
	auth := new(models.PermissionAuth)
	auth.Id = this.getInt64("id", 0)
	if _, err := auth.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
