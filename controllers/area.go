package controllers

import (
	"zd112_backstage/models"
	"time"
)

type AreaController struct {
	BaseController
}

func (this *AreaController) List() {
	this.pageTitle("区域管理")
	this.display("area/list")
}

func (this *AreaController) Add() {
	this.pageTitle("增加区域")
	this.display("area/add")
}

func (this *AreaController) Edit() {
	this.pageTitle("编辑区域")
	area := new(models.Area)
	area.Id = this.getId64(0)
	if err := area.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, area, true, "area/edit")
}

func (this *AreaController) AjaxSave() {
	area := new(models.Area)
	area.Id = this.getId64(0)
	area.ParentId = this.parentId
	area.Name = this.getString("name", "不能为空!", 1)
	area.NameAb = this.getString("nameAb", "", 0)
	area.Icon = this.getString("file", "", 0)
	area.NameCn = this.getString("nameCn", "", 0)
	area.NameEn = this.getString("nameEn", "", 0)
	area.ZoneCode = this.getInt("zoneCode", 0)
	area.ZipCode = this.getInt("zipCode", 0)
	area.AreaCode = this.getString("areaCode", "", 0)
	area.TimeDifference = this.getString("timeDifference", "", 0)
	var err error
	if area.Id == 0 {
		area.CreateId = this.userId
		area.CreateTime = time.Now().Unix()
		_, err = area.Add()
	} else {
		area.UpdateId = this.userId
		area.UpdateTime = time.Now().Unix()
		_, err = area.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AreaController) Table() {
	area := new(models.Area)
	result, count := area.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AreaController) AjaxDel() {
	area := new(models.Area)
	area.Id = this.getId64(0)
	if _, err := area.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
