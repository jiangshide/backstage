package controllers

type IndexController struct {
	BaseWebController
}

func (this *IndexController) Index() {
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "index.html"
}

func (this *IndexController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display("home/start")
}
