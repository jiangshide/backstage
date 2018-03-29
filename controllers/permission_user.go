package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"zd112_backstage/models"
	"github.com/jiangshide/GoComm/utils"
	"runtime"
)

type PermissionUserController struct {
	BaseController
}

func (this *PermissionUserController) List() {
	this.pageTitle("用户列表")
	this.display(this.getBgPermissionAction("user/list"))
}

func (this *PermissionUserController) Add() {
	this.pageTitle("增加用户")
	this.display(this.getBgPermissionAction("user/add"))
}

func (this *PermissionUserController) Edit() {
	this.pageTitle("编辑用户")
	user := new(models.PermissionUser)
	user.Id = this.getId64(0)

	if err := user.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, user, true, this.getBgPermissionAction("user/edit"))
}

func (this *PermissionUserController) AjaxSave() {
	user := new(models.PermissionUser)
	user.Id = this.getId64(0)
	user.Name = this.getString("name", "账号不能为空!", defaultMinSize)
	user.Salt = utils.GetRandomString(10)
	password := this.getString("password", "密码不能为空!", defaultMinSize)
	user.Password = utils.Md5(password + user.Salt)
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

func (this *PermissionUserController) Table() {
	user := new(models.PermissionUser)
	result, count := user.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *PermissionUserController) AjaxDel() {
	user := new(models.PermissionUser)
	user.Id = this.getId64(0)
	if _, err := user.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *PermissionUserController) Reg() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		user := new(models.PermissionUser)
		user.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		repassword := this.getString("repassword", "请再次输入密码!", defaultMinSize)
		if password != repassword {
			this.ajaxMsg("密码不一致!", MSG_ERR)
		}
		user.Salt = utils.GetRandomString(10)
		user.Password = utils.Md5(password + user.Salt)
		user.CreateTime = time.Now().Unix()
		if _, err := user.Add(); err != nil {
			this.showTips(err)
		}
		this.redirect(beego.URLFor("UserController.Login"))
	}
	this.TplName = "user/reg.html"
}

func (this *PermissionUserController) Login() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		user := new(models.PermissionUser)
		user.Name = this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		if err := user.Query(); err != nil {
			this.showTips("该账号不存在!")
		}
		if user.Status == 0 {
			this.showTips("该账号未激活!")
		} else if user.Status == 2 {
			this.showTips("该账号已被禁用!")
		} else if user.Password != utils.Md5(password+user.Salt) {
			this.showTips("该账号密码错误!")
		}
		user.UpdateTime = time.Now().Unix()
		if _, err := user.Update(); err == nil {
			userLocation := new(models.UserLocation)
			userLocation.UserId = user.Id
			this.userId = user.Id
			this.userName = user.Name
			userLocation.Ip = this.getIp()
			userLocation.Mac = this.getMac()
			userLocation.Device = runtime.GOOS
			userLocation.Arch = runtime.GOARCH
			userLocation.SdkVersion = runtime.Version()
			userLocation.AppVersion = this.version
			userLocation.CreateId = this.userId
			userLocation.CreateTime = time.Now().Unix()
			if index, err := userLocation.Add(); err != nil {
				beego.Info(index, " | err:", err)
			}
			this.setCook(user, 10000)
			this.redirect("/")
		} else {
			this.showTips(err)
		}
	}
	this.display(this.getBgPermissionAction("user/login"))
}

func (this *PermissionUserController) LogOut() {
	this.Ctx.SetCookie(AUTH, "")
	this.redirect(beego.URLFor("UserController.Login"))
}

func (this *PermissionUserController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}
