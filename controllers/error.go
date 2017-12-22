package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (this *ErrorController)Error404(){
	this.Data["content"]="正在开发中..."
	this.TplName="error/404.html"
}

func (this *ErrorController)Error501(){
	this.Data["content"]="server error"
	this.TplName = "error/501.html"
}

func (this *ErrorController)ErrorDb(){
	this.Data["content"]="database is now down"
	this.TplName = "error/dberror.html"
}