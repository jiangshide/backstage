package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"zd112_backstage/models"
)

type AreaCodeController struct {
	BaseController
}

func (this *AreaCodeController) List() {
	this.pageTitle("民族列表")
	areaCode := new(models.AreaCode)
	if _, count := areaCode.List(this.pageSize, this.offSet); count == 0 {
		beego.Info("-------count:", count)
		this.parseExcel()
	}
	this.display("nation/list")
}

func (this *AreaCodeController) Add() {
	this.pageTitle("新增民族名称")
	this.display("area_code/add")
}

func (this *AreaCodeController) Edit() {
	this.pageTitle("编辑民族名称")
	areaCode := new(models.AreaCode)
	areaCode.Id = this.getId64(0)
	if err := areaCode.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, areaCode, true, "area_code/edit")
}

func (this *AreaCodeController) AjaxSave() {
	areaCode := new(models.AreaCode)
	areaCode.Id = this.getId64(0)
	areaCode.NameCn = this.getString("nameCn", "名称不能为空!", 1)
	areaCode.NameEn = this.getString("nameEn", "名称不能为空!", 1)
	var err error
	if areaCode.Id == 0 {
		areaCode.CreateId = this.userId
		areaCode.CreateTime = time.Now().Unix()
		_, err = areaCode.Add()
	} else {
		areaCode.UpdateId = this.userId
		areaCode.UpdateTime = time.Now().Unix()
		_, err = areaCode.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AreaCodeController) Table() {
	areaCode := new(models.AreaCode)
	result, count := areaCode.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AreaCodeController) AjaxDel() {
	areaCode := new(models.AreaCode)
	areaCode.Id = this.getId64(0)
	if _, err := areaCode.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AreaCodeController) parseExcel(){

}