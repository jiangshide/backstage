package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"zd112_backstage/utils"
	"zd112_backstage/models"
	"github.com/george518/PPGo_ApiAdmin/libs"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) List() {
	this.Data["pageTitle"] = "管理员管理"
	this.display()
	//self.TplName = "admin/list.html"
}

func (this *AdminController) Add() {
	this.Data["pageTitle"] = "新增管理员"

	// 角色
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	role := new(models.Role)
	result, _ := role.GetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}

	this.Data["role"] = list

	this.display()
}

func (this *AdminController) Edit() {
	this.Data["pageTitle"] = "编辑管理员"

	user := new(models.User)
	user.Id = this.getId()
	user.Query()
	this.Data["user"] = user
	profile := new(models.Profile)
	profile.Id = user.Id;
	profile.Query();
	this.Data["profile"] = profile

	roleIds := strings.Split(user.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	role := new(models.Role)
	result, _ := role.GetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(roleIds); i++ {
			role_id, _ := strconv.ParseInt(roleIds[i], 10, 64)
			if role_id == v.Id {
				row["checked"] = 1
			}
			fmt.Println(roleIds[i])
		}
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}
	this.Data["role"] = list
	this.display()
}

func (this *AdminController) AjaxSave() {
	id := this.getId()
	if id == 0 {
		Admin := new(models.User)
		Admin.Name = this.getString("name")
		Admin.RoleIds = this.getString("roleids")
		Admin.UpdateTime = time.Now().Unix()
		Admin.UpdateId = this.userId
		Admin.Status = 1

		// 检查登录名是否已经存在
		_, err := Admin.GetByName(Admin.Name)

		if err == nil {
			this.ajaxMsg("登录名已经存在", MSG_ERR)
		}
		//新增
		pwd, salt := utils.Password(4, "")
		Admin.Password = pwd
		Admin.Salt = salt
		Admin.CreateTime = time.Now().Unix()
		Admin.CreateId = this.userId
		if _, err := models.Add(Admin); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	admin := new(models.User)
	admin.Id = id
	if err := admin.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	admin.Name = this.getString("name")
	admin.RoleIds = this.getString("roleids")
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	admin.Status = 1

	resetPwd, _ := this.GetInt("reset_pwd")
	if resetPwd == 1 {
		pwd, salt := libs.Password(4, "")
		admin.Password = pwd
		admin.Salt = salt
	}
	if _, err := admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg(strconv.Itoa(resetPwd), MSG_OK)
}

func (this *AdminController) AjaxStatus() {

	Admin_id := this.getId()
	status := this.getString("status")
	if Admin_id == 1 {
		this.ajaxMsg("超级管理员不允许操作", MSG_ERR)
	}

	Admin_status := 0
	if status == "enable" {
		Admin_status = 1
	}
	admin := new(models.User)
	Admin, _ := admin.GetById(Admin_id)
	Admin.UpdateTime = time.Now().Unix()
	Admin.Status = Admin_status
	Admin.Id = Admin_id

	if _, err := Admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("操作成功", MSG_OK)
}

func (this *AdminController) AjaxDel() {
	admin := new(models.User)
	admin.Id = this.getId()
	if _, err := admin.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("操作成功", MSG_OK)
}

func (this *AdminController) Table() {
	//列表
	page := this.getInt("page", 1)
	limit := this.getInt("limit", 30)
	realName := this.getString("realName")

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	admin := new(models.User)
	result, count := admin.GetList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
