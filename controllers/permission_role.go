package controllers

import (
	"time"
	"zd112_backstage/models"
)

type PermissionRoleController struct {
	BaseController
}

func (this *PermissionRoleController) List() {
	this.pageTitle("角色管理")
	this.display(this.getBgPermissionAction("role/list"))
}

func (this *PermissionRoleController) Add() {
	this.Data["zTree"] = true
	this.pageTitle("增加角色")
	this.display(this.getBgPermissionAction("role/add"))
}

func (this *PermissionRoleController) Edit() {
	this.pageTitle("编辑角色")
	role := new(models.PermissionRole)
	role.Id = this.getId64(0)
	if err := role.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, role, true, this.getBgPermissionAction("role/edit"))
}

func (this *PermissionRoleController) AjaxSave() {
	role := new(models.PermissionRole)
	role.Name = this.getString("role_name", "角色名称不能为空!", 0)
	role.Detail = this.getString("detail", "", 0)
	role.Id = this.getId64(0)
	var err error
	if role.Id == 0 {
		role.CreateId = this.userId
		role.CreateTime = time.Now().Unix()
		_, err = role.Add()
	} else {
		role.UpdateId = this.userId
		role.UpdateTime = time.Now().Unix()
		_, err = role.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *PermissionRoleController) Table() {
	role := new(models.PermissionRole)
	result, count := role.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *PermissionRoleController) AjaxDel() {
	role := new(models.PermissionRole)
	role.Id = this.getInt64("id", 0)
	if _, err := role.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
