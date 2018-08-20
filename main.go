package main

import (
	_ "backstage/routers"
	"github.com/astaxie/beego"
	"backstage/models"
)

func main() {
	models.Init()
	beego.Run()
}

