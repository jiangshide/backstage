package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	beego.Info("----------indexss")
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "index.html"
}

func (this *IndexController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display("home/start")
}


func init(){
	beego.Info("-----------------index")
}