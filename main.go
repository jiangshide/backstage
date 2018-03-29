package main

import (
	_ "zd112_backstage/routers"
	"github.com/astaxie/beego"
	"zd112_backstage/models"
)

func main() {
	models.Init()
	beego.Run()
}

