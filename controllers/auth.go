package controllers

import (
	"zd112_backstage/models"
	"time"
	"github.com/astaxie/beego"
)

type AuthController struct {
	BaseWebController
}

func (this *AuthController) Index() {
	this.pageTitle("权限因子")
	this.display()
}

func (this *AuthController) List() {
	this.pageTitle("权限因子列表")
	this.Data["zTree"] = true
	this.display("auth/list")
}

func (this *AuthController) GetNodes() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AuthList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Pid
		row["name"] = v.Name
		row["open"] = true
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AuthController) GetNode() {
	auth := new(models.Auth)
	auth.Id = this.getInt64("id", 0)
	if err := auth.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	row["id"] = auth.Id
	row["pid"] = auth.Pid
	row["name"] = auth.Name
	row["sort"] = auth.Sort
	row["is_show"] = auth.IsShow
	row["icon"] = auth.Icon
	this.ajaxList("成功", MSG_OK, 0, row)
}

func (this *AuthController) AjaxSave() {
	auth := new(models.Auth)
	auth.UserId = this.userId
	beego.Info("-------------userId:",this.userId)
	auth.Pid = this.getInt("pid", 0)
	beego.Info("------------Pid:",auth.Pid)
	auth.Name = this.getString("name", "名称不能为空!", 1)
	beego.Info("-----------name:",auth.Name)
	auth.Url = this.getString("url", "访问Action不能为空!", 1)
	beego.Info("-----------url:",auth.Url)
	auth.Sort = this.getInt("sort", 0)
	auth.IsShow = this.getInt("is_show", 0)
	auth.Icon = this.getString("icon", "菜单图标不能为空!", 1)
	beego.Info("----------icon:",auth.Icon)
	auth.UpdateTime = time.Now().Unix()
	auth.Status = 1
	id := this.getInt64("id", 0)
	var err error
	beego.Info("-------------auth:",auth)
	if id == 0 {
		auth.CreateTime = time.Now().Unix()
		auth.CreateId = this.userId
		_, err = auth.Add()
	} else {
		auth.Id = id
		auth.UpdateId = this.userId
		auth.UpdateTime = time.Now().Unix()
		_, err = auth.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AuthController) AjaxDel() {
	auth := new(models.Auth)
	auth.Id = this.getInt64("id", 0)
	if _, err := auth.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
