package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	beego.Info("----------home:",this.pageSize)
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "home.html"
}

func (this *HomeController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display()
}
