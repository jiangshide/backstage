package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"pack/models"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) List() {
	this.Data["pageTitle"] = "角色管理"
	this.display()
}

func (this *RoleController) Add() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "新增角色"
	this.display()
}
func (this *RoleController) Edit() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "编辑角色"

	id := this.getId()
	role1 := new(models.Role)
	role, _ := role1.GetById(id)
	row := make(map[string]interface{})
	row["id"] = role.Id
	row["role_name"] = role.RoleName
	row["detail"] = role.Detail
	this.Data["role"] = row

	//获取选择的树节点
	roleAuth1 := new(models.RoleAuth)
	roleAuth, _ := roleAuth1.GetById(id)
	authId := make([]int64, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	this.Data["auth"] = authId
	fmt.Println(authId)
	this.display()
}

func (this *RoleController) AjaxSave() {
	role := new(models.Role)
	role.RoleName = strings.TrimSpace(this.GetString("role_name"))
	role.Detail = strings.TrimSpace(this.GetString("detail"))
	role.CreateTime = time.Now().Unix()
	role.UpdateTime = time.Now().Unix()
	role.Status = 1
	auths := strings.TrimSpace(this.GetString("nodes_data"))
	role_id, _ := this.GetInt64("id")
	if role_id == 0 {
		//新增
		role.CreateTime = time.Now().Unix()
		role.UpdateTime = time.Now().Unix()
		role.CreateId = this.userId
		role.UpdateId = this.userId
		if id, err := role.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			ra := new(models.RoleAuth)
			authsSlice := strings.Split(auths, ",")
			for _, v := range authsSlice {
				aid, _ := strconv.ParseInt(v,10,64)
				ra.AuthId = aid
				ra.RoleId = id
				ra.Add()
			}
		}
		this.ajaxMsg("", MSG_OK)
	}
	//修改
	role.Id = role_id
	role.UpdateTime = time.Now().Unix()
	role.UpdateId = this.userId
	if err := role.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		// 删除该角色权限
		//models.RoleAuthDelete(role_id)
		ra := new(models.RoleAuth)
		ra.Delete(role_id)
		authsSlice := strings.Split(auths, ",")
		for _, v := range authsSlice {
			aid, _ := strconv.ParseInt(v,10,64)
			ra.AuthId = aid
			ra.RoleId = int64(role_id)
			ra.Add()
		}

	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RoleController) AjaxDel() {

	role_id := this.getId()
	role1 := new(models.Role)
	role, _ := role1.GetById(role_id)
	role.Status = 0
	role.Id = role_id
	role.UpdateTime = time.Now().Unix()

	if _,err := role.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	// 删除该角色权限
	//models.RoleAuthDelete(role_id)
	this.ajaxMsg("", MSG_OK)
}

func (this *RoleController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	roleName := strings.TrimSpace(this.GetString("roleName"))
	this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if roleName != "" {
		filters = append(filters, "role_name__icontains", roleName)
	}
	role := new(models.Role)
	result, count := role.GetList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}