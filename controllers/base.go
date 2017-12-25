package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"fmt"
	"zd112_backstage/models"
	"time"
	"github.com/jiangshide/GoComm/utils"
	"github.com/jiangshide/tinify-go/tinify"
	"strconv"
	"reflect"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
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

func init() {
	beego.Info("--------------base--init")
}

var defaultTips = "该项不能为空!"
var defaultMinSize = 6

func (this *BaseController) prepare() {
	this.page, _ = this.GetInt("page", 1)
	this.pageSize, _ = this.GetInt("limit", 30)
	this.offSet = (this.page - 1) * this.pageSize
	Tinify.SetKey("nDuD_n78YCCvgJ7F5CZ_gvbrpU4iRtoZ")

	this.currParam()
}

func (this *BaseController) currParam() {
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	this.userIcon = "/static/mingzu/img/3.jpg"
	this.upload = "/static/upload/"
	this.Data["route"] = this.controller + "." + this.action
	this.Data["action"] = this.action
	this.defaultPsw = beego.AppConfig.String("defaultPsw")
	this.version = beego.AppConfig.String("version")

	this.auth()
	this.Data["userId"] = this.userId
	this.Data["userName"] = this.userName
	this.Data["userIcon"] = this.userIcon
	this.Data["isLogin"] = this.isLogin
	this.Data["logo"] = "/static/mingzu/img/11.jpg"
	this.Data["version"] = this.version
	this.Data["siteName"] = beego.AppConfig.String("site.app_name")
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, err := this.GetInt(key, defaultValue)
	if err != nil {
		beego.Info("------------key:", key, " | err:", err)
		this.showTips(err)
	}
	return resInt
}

func (this *BaseController) getInt64(key string, defaultValue int64) int64 {
	resInt, err := this.GetInt64(key, defaultValue)
	if err != nil {
		beego.Info("--------key:", key, " | err:", err)
		this.showTips(err)
	}
	return resInt
}

func (this *BaseController) getId(defaultValue int) int {
	return this.getInt("id", defaultMinSize)
}

func (this *BaseController) getId64(defaultValue int64) int64 {
	return this.getInt64("id", defaultValue)
}

func (this *BaseController) getGroupId(defaultValue int) int {
	return this.getInt("groupId", defaultValue)
}

func (this *BaseController) getGroupId64(defaultValue int64) int64 {
	return this.getInt64("groupId", defaultValue)
}

func (this *BaseController) showTips(errorMsg interface{}) {
	flash := beego.NewFlash()
	flash.Error(fmt.Sprint(errorMsg))
	flash.Store(&this.Controller)
	controller, action := this.GetControllerAndAction()
	this.Redirect(beego.URLFor(controller+"."+action), 302)
	this.StopRun()
}

func (this *BaseController) getString(key, tips string, minSize int) string {
	value := strings.TrimSpace(this.GetString(key, ""))
	errorMsg := ""
	if len(value) == 0 {
		errorMsg = tips
	} else if len(value) < minSize {
		errorMsg = "长度不能小于" + fmt.Sprint(minSize) + "字符"
	}
	if errorMsg != "" {
		beego.Info("--------------key:", key, " | errorMsg:", errorMsg)
		this.showTips(errorMsg)
	}
	return value
}

func (this *BaseController) pageTitle(pageTitle string) {
	this.Data["pageTitle"] = pageTitle
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) display(tpl ...string) {
	this.currParam()
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = this.controller + "/" + this.action + ".html"
	}
	this.Layout = "comm/layout.html"
	this.TplName = tplName
}

func (this *BaseController) getBgUserAction(action string) string {
	return "user/" + action
}

func (this *BaseController) getBgAreaAction(action string) string {
	return "area/" + action
}

func (this *BaseController) getBgWebAction(action string) string {
	return "web/" + action
}

func (this *BaseController) getBgToolAction(action string) string {
	return "tools/" + action
}

func (this *BaseController) getBgAppAction(action string) string {
	return this.getBgTestAction("app/" + action)
}

func (this *BaseController) getBgTestAction(action string) string {
	return "test/" + action
}

func (this *BaseController) getBgApiAction(action string) string {
	return "api/" + action
}

func (this *BaseController) ajaxMsg(msg interface{}, msgNo int) {
	out := make(map[string]interface{})
	out["status"] = msgNo
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxList(msg interface{}, msgNo int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgNo
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.Data["data"] = data
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) row(row map[string]interface{}, model interface{}, isEdit bool) {
	this.Data["row"] = this.relectRow(row, model, 0, isEdit)
}

func (this *BaseController) parse(list []map[string]interface{}, row map[string]interface{}, k int, v interface{}, isEdit bool) {
	if list == nil {
		return
	}
	list[k] = this.relectRow(row, v, 0, isEdit)
}

func (this *BaseController) group(list []map[string]interface{}, row map[string]interface{}, k int, v interface{}, id int64, isEdit bool) {
	if list == nil {
		return
	}
	list[k] = this.relectRow(row, v, id, isEdit)
}

func (this *BaseController) relectRow(row map[string]interface{}, v interface{}, id int64, isEdit bool) map[string]interface{} {
	var field reflect.Type
	var value reflect.Value
	field = reflect.TypeOf(v).Elem()
	value = reflect.ValueOf(v).Elem()
	size := field.NumField()
	var fieldName string
	if row == nil {
		row = make(map[string]interface{})
	}
	for i := 0; i < size; i++ {
		v := value.Field(i)
		fieldName = field.Field(i).Name
		if fieldName == "Id" && id > 0 {
			if v.Int() == id {
				row["Selected"] = true
			} else {
				row["Selected"] = false
			}
			row[fieldName] = v.Int()
			continue
		}
		switch value.Field(i).Kind() {
		case reflect.Bool:
			row[fieldName] = v.Bool()
		case reflect.Int:
			if v.Int() > 0 {
				row[fieldName] = v.Int()
			}
		case reflect.Int64:
			if fieldName == "CreateTime" {
				if v.Int() > 0 && isEdit {
					row[fieldName] = v.Int()
				} else {
					row["CreateTimeFormat"] = beego.Date(time.Unix(v.Int(), 0), "Y-m-d H:i:s")
				}
				continue
			}
			if fieldName == "UpdateTime" {
				if v.Int() > 0 && !isEdit {
					row["UpdateTimeFormat"] = beego.Date(time.Unix(v.Int(), 0), "Y-m-d H:i:s")
				}
				continue
			}
			if fieldName == "CreateId" || fieldName == "UpdateId" {
				if v.Int() > 0 {
					if isEdit {
						if fieldName == "CreateId" {
							row[fieldName] = v.Int()
						}
					} else {
						admin := new(models.Admin)
						admin.Id = v.Int()
						if err := admin.Query(); err == nil {
							if fieldName == "CreateId" {
								row["CreateName"] = admin.Name
							} else {
								row["UpdateName"] = admin.Name
							}
						}
						row[fieldName] = v.Int()
					}
				}
				continue
			}
			row[fieldName] = v.Int()
		case reflect.String:
			if strings.Contains(v.String(), "static/") {
				row["File"] = v.String()
				this.setFileSize(row, v.String())
			} else {
				row[fieldName] = v.String()
			}
		}
	}
	beego.Info("-------row:", row)
	return row
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

func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	beego.Info("--------------------arr:", arr)
	this.userId = 0
	if len(arr) == 2 {
		idstr, psw := arr[0], arr[1]
		userId, _ := strconv.ParseInt(idstr, 11, 64)
		if userId > 0 {
			admin := new(models.Admin)
			admin.Id = userId
			err := admin.Query()
			if err == nil && psw == utils.Md5(this.getClientIp()+"|"+admin.Password+admin.Salt) {
				this.userId = admin.Id
				this.loginName = admin.Name
				this.userName = admin.RealName
				this.userIcon = "/static/mingzu/img/1.jpg"
				this.isLogin = true
				this.user = admin
			} else {
				beego.Error(err)
			}
			if strings.Contains(this.controller, "backstage") {
				this.AdminAuth()
			}
			//isHashAuth := strings.Contains(this.allowUrl, this.controller+"/"+this.action)
			//noAuth := "ajaxsave/ajaxdel/table/loginin/getnodes/start/show/ajaxapisace"
			//isNoAuth := strings.Contains(noAuth, this.action)
			//if isHashAuth == false && isNoAuth == false {
			//	this.Ctx.WriteString("没有权限")
			//	this.ajaxMsg("没有权限", MSG_ERR)
			//	return
			//}
		}
	}
	//if this.userId == 0 && (this.controller != "login" && this.action != "loginin") {
	//	this.redirect(beego.URLFor("LoginController.LoginIn"))
	//}
}

func (this *BaseController) AdminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if this.userId != 1 {
		adminAuthIds, _ := models.RoleAuthGetByIds(this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthList(1, 100, filters...)
	list := make([]map[string]interface{}, len(result))
	lists := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.Url != " " || v.Url != "/" {
			allow_url += v.Url
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["Name"] = v.Name
			row["Url"] = v.Url
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["Name"] = v.Name
			row["Url"] = v.Url
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			lists[j] = row
			j++
		}
	}
	this.Data["SideMenu1"] = list[:i]
	this.Data["SideMenu2"] = lists[:j]
	this.allowUrl = allow_url + "/index"
}
