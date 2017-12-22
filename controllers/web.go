package controllers

import (
	"zd112_backstage/models"
	"time"
)

type BannerController struct {
	BaseWebController
}

func (this *BannerController) List() {
	this.pageTitle("焦点图列表")
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
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, banner, true)
	this.display(this.getBgWebAction("banner/edit"))
}

func (this *BannerController) AjaxSave() {
	banner := new(models.Banner)
	banner.Name = this.getString("name", "名称不能为空!", 1)
	banner.Link = this.getString("link", "", 0)
	banner.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	banner.Descript = this.getString("descript", "", 0)
	banner.Id = this.getId64(0)
	var err error
	if banner.Id == 0 {
		banner.CreateId = this.userId
		banner.CreateTime = time.Now().Unix()
		_, err = banner.Add()
	} else {

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
