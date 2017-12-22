package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"fmt"
	"zd112_backstage/models"
	"time"
	"github.com/jiangshide/GoComm/utils"
	"github.com/jiangshide/tinify-go/tinify"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	isApi      bool
	version    string
	controller string
	action     string
	userId     int64
	userName   string
	loginName  string
	userIcon   string
	isLogin    bool
	logo       string
	page       int
	pageSize   int
	offSet     int
	allowUrl   string
	user       *models.Admin
	defaultPsw string
	upload     string
}

func (this *BaseController) prepare() {
	this.page, _ = this.GetInt("page", 1)
	this.pageSize, _ = this.GetInt("limit", 30)
	this.offSet = (this.page - 1) * this.pageSize
	Tinify.SetKey("nDuD_n78YCCvgJ7F5CZ_gvbrpU4iRtoZ")
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getClientIp() string {
	return this.Ctx.Input.IP()
}

func (this *BaseController) Upload() {
	f, fh, err := this.GetFile("file")
	defer f.Close()
	fileName := fh.Filename
	sufix := "default"
	if strings.Contains(fh.Filename, ".") {
		sufix = fileName[strings.LastIndex(fileName, ".")+1:]
	}
	fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	toFilePath := this.upload + sufix + "/" + fileName
	var size, resize int64
	if err = this.SaveToFile("file", utils.GetCurrentDir(toFilePath)); err == nil && this.getFileType(fileName) == "图片" {
		size, resize = this.compress(toFilePath)
	}
	this.ajaxMsgFile(toFilePath, size, resize, MSG_OK)
}

func (this *BaseController) ajaxMsgFile(msg interface{}, size, resize int64, msgNo int) {
	out := make(map[string]interface{})
	out["status"] = msgNo
	out["message"] = msg
	out["size"] = size
	out["resize"] = resize
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) getFileFormat(name string) (int64, int64, string) {
	size, sufix := utils.FileSize(name)
	format := new(models.Format)
	format.Name = sufix
	if err := format.Query(); err == nil {
		formatType := new(models.FormatType)
		formatType.Id = format.ParentId
		formatType.Query()
		return formatType.Id, size, sufix
	}
	return 0, size, sufix
}

func (this *BaseController) getFileType(name string) string {
	format := new(models.Format)
	format.Name = name[strings.LastIndex(name, ".")+1:]
	format.Query();
	formatType := new(models.FormatType)
	formatType.Id = format.ParentId
	formatType.Query()
	return formatType.Name
}

func (this *BaseController) compress(path string) (int64, int64) {
	path = utils.GetCurrentDir(path)
	size, _ := utils.FileSize(path)
	src, err := Tinify.FromFile(path)
	var resize int64
	if err == nil {
		if err = src.ToFile(path); err == nil {
			res, _ := utils.FileSize(path)
			resize = res
		}
	}
	if err != nil {
		beego.Error("compress:", err)
	}
	return size, resize
}

func (this *BaseController) setFileSize(row map[string]interface{}, file string) {
	size, _ := utils.FileSize(file)
	row["Size"] = size
}
