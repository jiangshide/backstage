package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"zd112_backstage/models"
	"github.com/jiangshide/GoComm/utils"
)

type UserController struct {
	BaseController
}

func (this *UserController) List() {
	this.pageTitle("用户列表")
	this.display(this.getBgUserAction("list"))
}

func (this *UserController) Add() {
	this.pageTitle("增加用户")
	this.display(this.getBgUserAction("add"))
}

func (this *UserController) Edit() {
	this.pageTitle("编辑用户")
	user := new(models.User)
	user.Id = this.getId64(0)
	if err := user.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, user, true)
	this.display(this.getBgUserAction("edit"))
}

func (this *UserController) AjaxSave() {
	user := new(models.User)
	user.Id = this.getId64(0)
	user.Name = this.getString("name", "名称不能为空!", defaultMinSize)
	user.Name = this.getString("username", "账号不能为空!", defaultMinSize)
	user.RealName = this.getString("realname", "真实名称不能为空!", 1)
	user.Salt = utils.GetRandomString(10)
	user.LastIp = this.getClientIp()
	password := this.getString("password", "密码不能为空!", defaultMinSize)
	user.Password = utils.Md5(password + user.Salt)
	user.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
	user.RoleIds = "1"
	user.Phone = this.getString("phone", "电话号码不能为空!", defaultMinSize)
	user.Sex = this.getInt("radio", 0)
	user.Motto = this.getString("motto", "请君赠言...", defaultMinSize)
	var err error
	if user.Id == 0 {
		user.CreateId = this.userId
		user.CreateTime = time.Now().Unix()
		_, err = user.Add()
	} else {
		user.UpdateId = this.userId
		user.UpdateTime = time.Now().Unix()
		_, err = user.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *UserController) Table() {
	list := make([]map[string]interface{}, 0)
	user := new(models.User)
	result, count := user.List(this.pageSize, this.offSet)
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *UserController) AjaxDel() {
	user := new(models.User)
	user.Id = this.getId64(0)
	if _, err := user.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *UserController) Reg() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		user := new(models.User)
		user.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		user.Password = password
		user.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
		user.Sex = this.getInt("radio", 0)
		user.LastIp = this.getClientIp()
		user.Salt = utils.GetRandomString(10)
		user.Password = utils.Md5(password + user.Salt)
		user.RoleIds = "2"
		user.CreateTime = time.Now().Unix()
		if _, err := user.Add(); err != nil {
			this.showTips(err)
		}
		this.redirect(beego.URLFor("UserController.Login"))
	}
	this.TplName = "login/reg.html"
}

func (this *UserController) Login() {
	if this.userId > 0 {
		this.redirect(beego.URLFor("MainController.Index"))
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		user := new(models.User)
		user.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		err := user.Query()
		if err != nil || user.Password != utils.Md5(password+user.Salt) {
			this.showTips(err)
		} else if user.Status == -1 {
			this.showTips("该账号已禁用!")
		}
		user.LastIp = this.getClientIp()
		user.LastLogin = time.Now().Unix()
		user.Update()
		authKey := utils.Md5(this.getClientIp() + "|" + user.Password + user.Salt)
		this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authKey, 7*86400)
		this.redirect(beego.URLFor("MainController.Index"))
	}
	this.TplName = "login/login.html"
}

func (this *UserController) LoginOut() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("IndexController.Index"))
}

func (this *UserController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}
