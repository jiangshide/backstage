package main

import (
	_ "pack/routers"
	"github.com/astaxie/beego"
	"pack/models"
)

func main() {
	models.Init()
	beego.Run()
}

