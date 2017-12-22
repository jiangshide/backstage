package controllers

import (
	"zd112_backstage/models"
	"time"
	"github.com/jiangshide/GoComm/utils"
)

type KeyController struct {
	BaseWebController
}

func (this *KeyController) List() {
	this.pageTitle("key列表")
	this.display(this.getBgApiAction("key/list"))
}

func (this *KeyController) Add() {
	this.pageTitle("key申请")
	this.display(this.getBgApiAction("key/add"))
}

func (this *KeyController) Edit() {
	this.pageTitle("编辑key")
	key := new(models.Key)
	key.Id = this.getId64(0)
	if err := key.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, key, true)
	this.display(this.getBgApiAction("key/edit"))
}

func (this *KeyController) AjaxSave() {
	key := new(models.Key)
	key.Id = this.getId64(0)
	key.Name = this.getString("name", "名称不能为空!", 3)
	key.Descript = this.getString("descript", "", 0)
	key.Key = utils.Md5(this.userName + key.Name)
	var err error
	if key.Id == 0 {
		key.CreateId = this.userId
		key.CreateTime = time.Now().Unix()
		_, err = key.Add()
	} else {
		key.UpdateId = this.userId
		key.UpdateTime = time.Now().Unix()
		_, err = key.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *KeyController) Table() {
	key := new(models.Key)
	result, count := key.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *KeyController) AjaxDel() {
	key := new(models.Key)
	key.Id = this.getId64(0)
	if _, err := key.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
