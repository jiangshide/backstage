package controllers

import (
	"zd112_backstage/models"
	"os"
	"github.com/astaxie/beego"
	"time"
	"bufio"
	"io"
	"strings"
)

type QrcodeController struct {
	BaseWebController
}

func (this *QrcodeController) List() {
	this.pageTitle("二维码列表")
	this.display(this.getBgToolAction("qrcode/list"))
}

func (this *QrcodeController) Add() {
	this.pageTitle("增加二维码")
	this.display(this.getBgToolAction("qrcode/add"))
}

func (this *QrcodeController) Edit() {
	this.pageTitle("编辑二维码")
	qrcode := new(models.Qrcode)
	qrcode.Id = this.getId64(0)
	if err := qrcode.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, qrcode,true)
	this.display(this.getBgToolAction("qrcode/edit"))
}

func (this *QrcodeController) AjaxSave() {
	qrcode := new(models.Qrcode)
	qrcode.Id = this.getId64(0)
	qrcode.Name = this.getString("name", "名称不能为空!", 1)
	qrcode.Descript = this.getString("descript", "", 0)
	var err error
	if qrcode.Id == 0 {
		qrcode.CreateId = this.userId
		qrcode.CreateTime = time.Now().Unix()
		_, err = qrcode.Add()
	} else {
		qrcode.CreateId = this.getInt64("create_id", 0)
		qrcode.CreateTime = this.getInt64("create_time", 0)
		qrcode.UpdateId = this.userId
		qrcode.UpdateTime = time.Now().Unix()
		_, err = qrcode.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *QrcodeController) Table() {
	qrcode := new(models.Qrcode)
	result, count := qrcode.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *QrcodeController) AjaxDel() {
	qrcode := new(models.Qrcode)
	qrcode.Id = this.getId64(0)
	if _, err := qrcode.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type FormatTypeController struct {
	BaseWebController
}

func (this *FormatTypeController) List() {
	this.pageTitle("文件类型列表")
	this.display(this.getBgToolAction("format/type/list"))
}

func (this *FormatTypeController) Add() {
	this.pageTitle("增加文件类型")
	this.display(this.getBgToolAction("format/type/add"))
}

func (this *FormatTypeController) Edit() {
	this.pageTitle("编辑文件类型")
	formatType := new(models.FormatType)
	formatType.Id = this.getId64(0)
	if err := formatType.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, formatType,true)
	this.display(this.getBgToolAction("format/type/edit"))
}

func (this *FormatTypeController) AjaxSave() {
	formatType := new(models.FormatType)
	formatType.Id = this.getId64(0)
	formatType.Name = this.getString("name", "名称不能为空!", 1)
	var err error
	if formatType.Id == 0 {
		formatType.CreateId = this.userId
		formatType.CreateTime = time.Now().Unix()
		_, err = formatType.Add()
	} else {
		formatType.UpdateId = this.userId
		formatType.UpdateTime = time.Now().Unix()
		_, err = formatType.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatTypeController) Table() {
	formatType := new(models.FormatType)
	result, count := formatType.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *FormatTypeController) AjaxDel() {
	formatType := new(models.FormatType)
	formatType.Id = this.getId64(0)
	if _, err := formatType.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type FormatController struct {
	BaseWebController
}

func (this *FormatController) List() {
	this.pageTitle("文件格式列表")
	this.display(this.getBgToolAction("format/list"))
}

func (this *FormatController) Add() {
	this.pageTitle("增加文件格式")
	this.parent(0)
	this.display(this.getBgToolAction("format/add"))
}

func (this *FormatController) Edit() {
	this.pageTitle("编辑文件格式")
	format := new(models.Format)
	format.Id = this.getId64(0)
	if err := format.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(format.ParentId)
	this.row(nil, format,true)
	this.display(this.getBgToolAction("format/edit"))
}

func (this *FormatController) AjaxSave() {
	format := new(models.Format)
	format.ParentId = this.getGroupId64(0)
	format.Name = this.getString("name", "名称不能为空!", 1)
	format.Id = this.getId64(0)
	var err error
	if format.Id == 0 {
		format.CreateId = this.userId
		format.UpdateTime = time.Now().Unix()
		_, err = format.Add()
	} else {
		format.UpdateId = this.userId
		format.UpdateTime = time.Now().Unix()
		_, err = format.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatController) parent(id int64) {
	formatType := new(models.FormatType)
	result, count := formatType.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *FormatController) Table() {
	format := new(models.Format)
	result, count := format.List(this.pageSize, this.offSet)
	beego.Info("result:", result, " | count:", count)
	if count == 0 {
		this.Export()
	}
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		formatType := new(models.FormatType)
		formatType.Id = v.ParentId
		if err := formatType.Query(); err == nil {
			row["Parent"] = formatType.Name
		} else {
			beego.Error(err)
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *FormatController) AjaxDel() {
	format := new(models.Format)
	format.Id = this.getId64(0)
	if _, err := format.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *FormatController) Export() {
	f, _ := os.Open("format")
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(line, ":") {
			lineStr := strings.Split(line, ":")
			if lineStr[0] == "video" {
				this.addDb("视频", lineStr[1])
			} else if lineStr[0] == "audio" {
				this.addDb("音频", lineStr[1])
			} else if lineStr[0] == "picture" {
				this.addDb("图片", lineStr[1])
			} else if lineStr[0] == "text" {
				this.addDb("文字", lineStr[1])
			}
		}
	}
}

func (this *FormatController) addDb(name, format string, ) {
	formatType := new(models.FormatType)
	formatType.Name = name
	formatType.CreateId = this.userId
	formatType.CreateTime = time.Now().Unix()
	if _, err := formatType.Add(); err != nil {
		beego.Error(err)
	}
	formatType.Query()
	if formatType.Id > 0 {
		formatStr := strings.Split(format, ".")
		for _, v := range formatStr {
			formatTable := new(models.Format)
			formatTable.ParentId = formatType.Id
			formatTable.Name = v;
			formatTable.CreateId = this.userId
			formatTable.CreateTime = time.Now().Unix()
			if _, err := formatTable.Add(); err != nil {
				beego.Error(err)
			}
		}
	}
}

type CompressController struct {
	BaseWebController
}

func (this *CompressController) List() {
	this.pageTitle("压缩文件列表")
	this.display(this.getBgToolAction("compress/list"))
}

func (this *CompressController) Add() {
	this.pageTitle("添加文件")
	this.display(this.getBgToolAction("compress/add"))
}

func (this *CompressController) Edit() {
	this.pageTitle("编辑文件")
	compress := new(models.Compress)
	compress.Id = this.getId64(0)
	if err := compress.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, compress,true)
	this.display(this.getBgToolAction("compress/edit"))
}

func (this *CompressController) AjaxSave() {
	compress := new(models.Compress)
	compress.Id = this.getId64(0)
	compress.Name = this.getString("name", "名称不能为空!", 1)
	compress.Url = this.getString("file", "上传文件URL为空!", defaultMinSize)
	compress.Descript = this.getString("descript", "", 0)
	compress.Size = this.getInt64("size", 0)
	compress.ReSize = this.getInt64("resize", 0)
	compress.Compress = this.getInt("compress", 0)
	if compress.ReSize > 0 {
		compress.Compress = 1
	}
	var err error
	if compress.Id == 0 {
		id, reSize, sufix := this.getFileFormat(compress.Url)
		compress.Type = id
		compress.Format = sufix
		compress.ReSize = reSize
		compress.CreateId = this.userId
		compress.CreateTime = time.Now().Unix()
		_, err = compress.Add()
	} else {
		compress.Downs = this.getInt("downs", 0)
		compress.UpdateId = this.userId
		compress.UpdateTime = time.Now().Unix()
		_, err = compress.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CompressController) Table() {
	compress := new(models.Compress)
	result, count := compress.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {

		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CompressController) AjaxDel() {
	compress := new(models.Compress)
	compress.Id = this.getId64(0)
	if _, err := compress.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
