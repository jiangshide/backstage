package controllers

import (
	"zd112_backstage/models"
	"time"
	"strconv"
	"strings"
	"github.com/jiangshide/GoComm/utils"
)

type AdminController struct {
	BaseWebController
}

func (this *AdminController) List() {
	this.pageTitle("管理员管理")
	this.display("backstage/admin/list")
}

func (this *AdminController) Add() {
	this.pageTitle("新增管理员")
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/add")
}

func (this *AdminController) Edit() {
	this.pageTitle("编辑管理员")
	admin := new(models.Admin)
	admin.Id = this.getInt64("id", 0)
	admin.Query()
	row := make(map[string]interface{})
	row["id"] = admin.Id
	row["login_name"] = admin.Name
	row["real_name"] = admin.RealName
	row["phone"] = admin.Phone
	row["email"] = admin.Email
	row["role_ids"] = admin.RoleIds
	this.Data["admin"] = row

	role_ids := strings.Split(admin.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(role_ids); i++ {
			role_id, _ := strconv.ParseInt(role_ids[i], 11, 64)
			if role_id == v.Id {
				row["checked"] = 1
			}
		}
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/edit")
}

func (this AdminController) AjaxSave() {
	admin := new(models.Admin)
	admin.Id = this.getInt64("id", 0)
	admin.Name = this.getString("login_name", "账号不能为空!", defaultMinSize)
	admin.RealName = this.getString("real_name", "真实名字不能为空!", defaultMinSize)
	admin.Phone = this.getString("phone", "联系电话不能为空!", defaultMinSize)
	admin.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
	admin.RoleIds = this.getString("roleids", "没有选择权限项!", 0)
	resetPsw := this.getInt("reset_pwd", 0)
	admin.Status = 1
	var err error
	if admin.Id == 0 {
		admin.Salt = utils.GetRandomString(10)
		admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
		admin.CreateTime = time.Now().Unix()
		admin.CreateId = this.userId
		_, err = admin.Add()
	} else {
		admin.UpdateTime = time.Now().Unix()
		admin.UpdateId = this.userId
		if resetPsw == 1 {
			admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
			admin.Salt = utils.GetRandomString(10)
		}
		_, err = admin.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg(strconv.Itoa(resetPsw), MSG_OK)
}

func (this AdminController) AjaxDel() {
	admin := new(models.Admin)
	admin.Id = this.getInt64("id", 0)
	if admin.Id == 1 {
		this.ajaxMsg("超级管理员不允许删除!", MSG_ERR)
	}
	if _, err := admin.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AdminController) Table() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	admin := new(models.Admin)
	result, count := admin.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
