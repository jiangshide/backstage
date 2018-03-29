package controllers

import (
	"time"
	"zd112_backstage/models"
	"github.com/astaxie/beego"
	"runtime"
)

type BannerController struct {
	BaseController
}

func (this *BannerController) List() {
	this.pageTitle("焦点图列表")

	runtime.GOMAXPROCS(runtime.NumCPU())
	go models.Cmd("/Users/etongdai/developer/client/react/reactnative/workspace/app/imgs/")
	beego.Info("---------------finish!")

	this.display(this.getBgWebAction("banner/list"))
}

func (this *BannerController) Add() {
	this.pageTitle("增加焦点图名称")
	this.display(this.getBgWebAction("banner/add"))
}

func (this *BannerController) Table() {
	banner := new(models.Banner)
	result, count := banner.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *BannerController) Edit() {
	this.pageTitle("编辑焦点图名称")
	banner := new(models.Banner)
	banner.Id = this.getId64(0)
	if err := banner.Query(); err != nil {
		beego.Info("--------err:", err)
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, banner, true, this.getBgWebAction("banner/edit"))
}

func (this *BannerController) AjaxSave() {
	banner := new(models.Banner)
	banner.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	banner.Title = this.getString("title", "专题名不能为空!", 1)
	banner.Name = this.getString("name", "名称不能为空!", 1)
	banner.Link = this.getString("link", "", 0)
	banner.Descript = this.getString("descript", "", 0)
	banner.Id = this.getId64(0)
	var err error
	if banner.Id == 0 {
		banner.CreateId = this.userId
		banner.CreateTime = time.Now().Unix()
		_, err = banner.Add()
	} else {
		banner.CreateId = this.getInt64("create_id", 0)
		banner.CreateTime = this.getInt64("create_time", 0)
		banner.UpdateId = this.userId
		banner.UpdateTime = time.Now().Unix()
		beego.Info("------update:", banner)
		_, err = banner.Update()
		beego.Info("---------err:", err)
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *BannerController) AjaxDel() {
	banner := new(models.Banner)
	banner.Id = this.getId64(0)
	if _, err := banner.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type NavController struct {
	BaseController
}

func (this *NavController) List() {
	this.pageTitle("导航栏列表")
	this.display(this.getBgWebAction("nav/list"))
}

func (this *NavController) Add() {
	this.pageTitle("增加导航栏名称")
	this.display(this.getBgWebAction("nav/add"))
}

func (this *NavController) Table() {
	nav := new(models.Nav)
	result, count := nav.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *NavController) Edit() {
	this.pageTitle("编辑导航名称")
	nav := new(models.Nav)
	nav.Id = this.getId64(0)
	if err := nav.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, nav, true, this.getBgWebAction("nav/edit"))
}

func (this *NavController) AjaxSave() {
	nav := new(models.Nav)
	nav.Id = this.getId64(0)
	nav.Name = this.getString("name", "名称不能为空!", 2)
	nav.Action = this.getString("action", "触发事件不能为空!", 1)
	nav.Descript = this.getString("descript", "", 0)
	var err error
	if nav.Id == 0 {
		nav.CreateId = this.userId
		nav.CreateTime = time.Now().Unix()
		_, err = nav.Add()
	} else {
		nav.CreateId = this.getInt64("create_id", 0)
		nav.CreateTime = this.getInt64("create_time", 0)
		nav.UpdateId = this.userId
		nav.UpdateTime = time.Now().Unix()
		_, err = nav.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *NavController) AjaxDel() {
	nav := new(models.Nav)
	nav.Id = this.getId64(0)
	if _, err := nav.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type UniversityController struct {
	BaseController
}

func (this *UniversityController) List() {
	this.pageTitle("学府列表")
	this.parentIds()
	this.display(this.getBgWebAction("nav/university/list"))
}

func (this *UniversityController) Add() {
	this.pageTitle("增加学府名称")
	beego.Info("-------------parentId:", this.parentId, " | userId:", this.userId)
	this.display(this.getBgWebAction("nav/university/add"))
}

func (this *UniversityController) Table() {
	university := new(models.University)
	result, count := university.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *UniversityController) Edit() {
	this.pageTitle("编辑学府名称")
	university := new(models.University)
	university.Id = this.getId64(0)
	if err := university.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, university, true, this.getBgWebAction("nav/university/edit"))
}

func (this *UniversityController) AjaxSave() {
	university := new(models.University)
	university.Id = this.getId64(0)
	university.GroupId = this.parentId
	university.Name = this.getString("name", "名称不能为空!", 1)
	university.Action = this.getString("action", "触发事件不能为空!", 1)
	university.Descript = this.getString("descript", "", 0)
	var err error
	beego.Info("-------------parentId:", this.parentId, " | id:", university.Id)
	if university.Id == 0 {
		university.CreateId = this.userId
		university.CreateTime = time.Now().Unix()
		_, err = university.Add()
	} else {
		university.CreateId = this.getInt64("create_id", 0)
		university.UpdateTime = this.getInt64("create_time", 0)
		university.UpdateId = this.userId
		university.UpdateTime = time.Now().Unix()
		_, err = university.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *UniversityController) AjaxDel() {
	university := new(models.University)
	university.Id = this.getId64(0)
	if _, err := university.Delete(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
