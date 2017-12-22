package controllers

import (
	"zd112_backstage/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseWebController
}

func (this *IndexController) Get() {
	this.Data["img"] = "/static/mingzu/img/1.jpg"
	banner := new(models.Banner)
	result, _ := banner.List(this.pageSize, this.offSet)
	for _,v := range result{
		beego.Info("---------name:",v.Icon)
	}
	this.Data["row"] = result
	this.display("index")
}
