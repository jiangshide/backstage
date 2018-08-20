package controllers

import (
	"strings"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jiangshide/GoComm/utils"
)

type UploadController struct{
	BaseController
}

func (this *UploadController) Upload() {
	beego.Info("-------------upload")
	f, fh, err := this.GetFile("file")
	defer f.Close()
	fileName := fh.Filename
	sufix := "default"
	if strings.Contains(fh.Filename, ".") {
		sufix = fileName[strings.LastIndex(fileName, ".")+1:]
	}
	fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	toFilePath := this.upload + sufix + "/" + fileName
	beego.Info("----------upload:", this.upload, " | sufix:", sufix, " | fileName:", fileName, " | toFilePath:", toFilePath)
	var size, resize int64
	path := utils.GetCurrentDir(toFilePath)
	beego.Info("-----------path:", path)
	if err = this.SaveToFile("file", path); err == nil {
		beego.Info("------err:", err)
		this.ajaxMsgFile(toFilePath, size, resize, MSG_OK)
	} else {
		beego.Info("------err:", err)
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
}