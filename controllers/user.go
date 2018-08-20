package controllers

import (
	"strings"
	"time"

	"backstage/models"
	"backstage/utils"
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Edit() {
	this.Data["pageTitle"] = "资料修改"
	id := this.userId
	admin := new(models.User)
	Admin, _ := admin.GetById(id)
	row := make(map[string]interface{})
	row["id"] = Admin.Id
	row["name"] = Admin.Name
	this.Data["admin"] = row
	this.display()
}

func (this *UserController) AjaxSave() {
	Admin_id, _ := this.GetInt64("id")
	admin := new(models.User)
	Admin, _ := admin.GetById(Admin_id)
	//修改
	Admin.Id = Admin_id
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = this.userId
	Admin.Name = this.getString("name")

	resetPwd := this.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(this.GetString("password_old"))
		pwdOldMd5 := utils.Md5([]byte(pwdOld + Admin.Salt))
		if Admin.Password != pwdOldMd5 {
			this.ajaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(this.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(this.GetString("password_new2"))

		if pwdNew1 != pwdNew2 {
			this.ajaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := utils.Password(4, pwdNew1)
		Admin.Password = pwd
		Admin.Salt = salt
	}
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = this.userId
	Admin.Status = 1

	if _,err := Admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

//登录 TODO:XSRF过滤
func (this *UserController) Login() {
	beego.Info("--------------------login~userId:",this.userId)
	if this.userId > 0 {
		this.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		username := this.getString("username")
		password := this.getString("password")

		if username != "" && password != "" {
			admin := new(models.User)
			user, err := admin.GetByName(username)
			fmt.Println(user)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != utils.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authkey := utils.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id,10)+"|"+authkey, 7*86400)

				this.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("UserController.Login"))
		}
	}
	this.TplName = "login/login.html"
}

//登出
func (self *UserController) Quit() {
	beego.Info("----------------quit")
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("UserController.Login"))
}

func (self *UserController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}