package main

import (
	"github.com/astaxie/beego"
	"zd112_backstage/models"
)

func main(){
	models.Init()
	beego.Run()
}