package controllers

import (
	"zd112_backstage/models"
	"github.com/astaxie/beego"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"time"
	"github.com/jiangshide/GoComm/utils"
)

type NationController struct {
	BaseController
}

func (this *NationController) List() {
	this.pageTitle("民族列表")
	this.display("nation/list")
}

func (this *NationController) Add() {
	this.pageTitle("新增民族名称")
	this.display("nation/add")
}

func (this *NationController) Edit() {
	this.pageTitle("编辑民族名称")
	nation := new(models.Nation)
	nation.Id = this.getId64(0)
	if err := nation.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, nation,true)
	this.display("nation/edit")
}

func (this *NationController) AjaxSave() {
	beego.Info("-----nation-----form:",this.Ctx.Request.Form)
	nation := new(models.Nation)
	nation.Id = this.getId64(0)
	nation.Name = this.getString("name", "名称不能为空!", 1)
	nation.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	var err error
	if nation.Id == 0 {
		nation.CreateId = this.userId
		nation.CreateTime = time.Now().Unix()
		_, err = nation.Add()
	}else{
		nation.UpdateId = this.userId
		nation.UpdateTime = time.Now().Unix()
		_,err = nation.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *NationController) Table() {
	nation := new(models.Nation)
	result, count := nation.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *NationController) AjaxDel() {
	nation := new(models.Nation)
	nation.Id = this.getId64(0)
	if _, err := nation.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *NationController) download() {
	mingzu := [...]string{"阿昌族", "白族", "保安族", "布朗族", "布依族", "藏族", "朝鲜族", "达翰尔族", "傣族", "昂德族", "东乡族", "侗族", "独龙族", "俄罗斯族", "鄂伦春族", "鄂温克族", "高山族", "哈尼族", "哈萨克族", "汉族", "赫哲族", "回族", "基诺族", "京族",
		"景颇族", "柯尔克孜族", "拉祜族", "黎族", "傈僳族", "珞巴族", "满族", "毛南族", "门巴族", "蒙古族", "苗族", "仫佬族", "纳西族", "怒族", "普米族", "羌族", "撒拉族", "畲族", "水族", "塔吉克族", "塔塔尔族", "土家族", "图族", "佤族", "维吾尔族", "乌孜别克族", "锡伯族", "瑶族", "彝族", "仡佬族", "裕固族", "壮族"}
	for k, v := range mingzu {
		index := ""
		if k < 9 {
			index = "0" + fmt.Sprint(k+1)
		} else {
			index = fmt.Sprint(k + 1)
		}
		imgRes, _ := http.Get("http://www.sinobuy.cn/mzfz/img/0" + index + ".jpg")
		defer imgRes.Body.Close()
		imgByte, _ := ioutil.ReadAll(imgRes.Body)
		dir := "/static/upload/img/"
		fileName := utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + ".jpg"
		path := utils.GetCurrentDir(dir) + fileName
		_, fErr := os.Stat(path)
		var fh *os.File
		if fErr != nil {
			fh, _ = os.Create(path)
		} else {
			fh, _ = os.Open(path)
		}
		defer fh.Close()
		if _, err := fh.Write(imgByte); err == nil {
			nation := new(models.Nation)
			nation.Name = v
			nation.Icon = dir + fileName
			nation.CreateId = this.userId
			nation.CreateTime = time.Now().Unix()
			if _, err := nation.Add(); err != nil {
				beego.Error("---------add:name:", v, err)
			}
		} else {
			beego.Error("------------write:name:", v, err)
		}
	}
}
