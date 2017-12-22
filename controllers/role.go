package controllers

import (
	"time"
	"zd112_backstage/models"
	"github.com/astaxie/beego"
)

type RoleController struct {
	BaseWebController
}

func (this *RoleController) List() {
	this.pageTitle("角色管理")
	this.display("backstage/role/list")
}

func (this *RoleController) Add() {
	this.Data["zTree"] = true
	this.pageTitle("增加角色")
	this.display("backstage/role/add")
}

func (this *RoleController) Edit() {
	this.Data["zTree"] = true
	this.pageTitle("编辑角色")
	role := new(models.Role)
	id := this.getInt64("id", 0)
	role.Id = id
	if err := role.Query(); err == nil {
		this.Data["role"] = role
	} else {
		beego.Error(err)
	}
	roleAuth := new(models.RoleAuth)
	roleAuth.RoleId = id
	if roleAuthList, err := roleAuth.List(); err == nil {
		authId := make([]int64, 0)
		for _, v := range roleAuthList {
			authId = append(authId, v.AuthId)
		}
		this.Data["auth"] = authId
	} else {
		beego.Error(err)
	}
	this.display("backstage/role/edit")
}

func (this *RoleController) AjaxSave() {
	role := new(models.Role)
	role.Name = this.getString("role_name", "角色名称不能为空!", 0)
	role.Detail = this.getString("detail", "", -1)
	role.CreateTime = time.Now().Unix()
	role.Status = 1
}

func (this *RoleController) AjaxDel() {
	role := new(models.Role)
	role.Id = this.getInt64("id", 0)
	if _, err := role.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RoleController) Table() {
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	this.pageSize = limit
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.RoleList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
