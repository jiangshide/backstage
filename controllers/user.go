package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"zd112_backstage/models"
	"github.com/jiangshide/GoComm/utils"
)

type UserController struct {
	BaseWebController
}

func (this *UserController) Admin() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		admin := new(models.Admin)
		admin.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		admin.RealName = this.getString("realname", "真实名称不能为空!", 1)
		admin.Salt = utils.GetRandomString(10)
		admin.LastIp = this.getClientIp()
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		admin.Password = utils.Md5(password + admin.Salt)
		admin.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
		admin.RoleIds = "1"
		admin.Phone = this.getString("phone", "电话号码不能为空!", defaultMinSize)
		admin.Sex = this.getInt("radio", 0)
		admin.Motto = this.getString("motto", "请君赠言...", defaultMinSize)
		admin.CreateTime = time.Now().Unix()
		if _, err := admin.Add(); err != nil {
			this.showTips(err)
		}
		this.redirect(beego.URLFor("UserController.Login"))
	}
	this.TplName = "login/admin.html"
}

func (this *UserController) Reg() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		admin := new(models.Admin)
		admin.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		admin.Password = password
		admin.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
		admin.Sex = this.getInt("radio", 0)
		admin.LastIp = this.getClientIp()
		admin.Salt = utils.GetRandomString(10)
		admin.Password = utils.Md5(password + admin.Salt)
		admin.RoleIds = "2"
		admin.CreateTime = time.Now().Unix()
		if _, err := admin.Add(); err != nil {
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
		admin := new(models.Admin)
		admin.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		err := admin.Query()
		if err != nil || admin.Password != utils.Md5(password+admin.Salt) {
			this.showTips(err)
		} else if admin.Status == -1 {
			this.showTips("该账号已禁用!")
		}
		admin.LastIp = this.getClientIp()
		admin.LastLogin = time.Now().Unix()
		admin.Update()
		authKey := utils.Md5(this.getClientIp() + "|" + admin.Password + admin.Salt)
		this.Ctx.SetCookie("auth", strconv.FormatInt(admin.Id, 10)+"|"+authKey, 7*86400)
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

func (this *UserController) Edit() {
	this.Data["pageTitle"] = "资料修改"
	admin := new(models.Admin)
	admin.Id = this.userId
	if err := admin.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, admin,true)
	this.display("backstage/user/edit")
}
