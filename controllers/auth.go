package controllers

import (
	"fmt"
	"time"
	"pack/models"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Index() {
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

func (this *AuthController) List() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

//获取全部节点
func (this *AuthController) GetNodes() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Pid
		row["name"] = v.AuthName
		row["open"] = true
		list[k] = row
	}

	this.ajaxList("成功", MSG_OK, count, list)
}

//获取一个节点
func (this *AuthController) GetNode() {
	id := this.getId()
	result, _ := models.AuthGetById(id)
	// if err == nil {
	// 	self.ajaxMsg(err.Error(), MSG_ERR)
	// }
	row := make(map[string]interface{})
	row["id"] = result.Id
	row["pid"] = result.Pid
	row["auth_name"] = result.AuthName
	row["auth_url"] = result.AuthUrl
	row["sort"] = result.Sort
	row["is_show"] = result.IsShow
	row["icon"] = result.Icon

	fmt.Println(row)

	this.ajaxList("成功", MSG_OK, 0, row)
}

//新增或修改
func (this *AuthController) AjaxSave() {
	auth := new(models.Auth)
	auth.UserId = this.userId
	auth.Pid = this.getInt64("pid",0)
	auth.AuthName = this.getString("auth_name")
	auth.AuthUrl = this.getString("auth_url")
	auth.Sort = this.getInt("sort",0)
	auth.IsShow = this.getInt("is_show",0)
	auth.Icon = this.getString("icon")
	auth.UpdateTime = time.Now().Unix()

	auth.Status = 1

	id := this.getId()
	if id == 0 {
		//新增
		auth.CreateTime = time.Now().Unix()
		auth.CreateId = this.userId
		auth.UpdateId = this.userId
		if _, err := models.AuthAdd(auth); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		auth.Id = id
		auth.UpdateId = this.userId
		if err := auth.Update(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	}

	this.ajaxMsg("", MSG_OK)
}

//删除
func (this *AuthController) AjaxDel() {
	id := this.getId()
	auth, _ := models.AuthGetById(id)
	auth.Id = id
	auth.Status = 0
	if err := auth.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}