package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"zd112_backstage/models"
	"reflect"
	"time"
	"fmt"
	"github.com/jiangshide/GoComm/utils"
)

type BaseWebController struct {
	BaseController
}

var defaultTips = "该项不能为空!"
var defaultMinSize = 6

func (this *BaseWebController) Prepare() {
	this.isApi = false
	this.prepare()
	this.currParam()
}

func (this *BaseWebController) getInt(key string, defaultValue int) int {
	resInt, err := this.GetInt(key, defaultValue)
	if err != nil {
		beego.Info("------------key:", key, " | err:", err)
		this.showTips(err)
	}
	return resInt
}

func (this *BaseWebController) getInt64(key string, defaultValue int64) int64 {
	resInt, err := this.GetInt64(key, defaultValue)
	if err != nil {
		beego.Info("--------key:", key, " | err:", err)
		this.showTips(err)
	}
	return resInt
}

func (this *BaseWebController) getId(defaultValue int) int {
	return this.getInt("id", defaultMinSize)
}

func (this *BaseWebController) getId64(defaultValue int64) int64 {
	return this.getInt64("id", defaultValue)
}

func (this *BaseWebController) getGroupId(defaultValue int) int {
	return this.getInt("groupId", defaultValue)
}

func (this *BaseWebController) getGroupId64(defaultValue int64) int64 {
	return this.getInt64("groupId", defaultValue)
}

func (this *BaseWebController) showTips(errorMsg interface{}) {
	flash := beego.NewFlash()
	flash.Error(fmt.Sprint(errorMsg))
	flash.Store(&this.Controller)
	controller, action := this.GetControllerAndAction()
	this.Redirect(beego.URLFor(controller+"."+action), 302)
	this.StopRun()
}

func (this *BaseWebController) getString(key, tips string, minSize int) string {
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

func (this *BaseWebController) currParam() {
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

func (this *BaseWebController) pageTitle(pageTitle string) {
	this.Data["pageTitle"] = pageTitle
}

func (this *BaseWebController) auth() {
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

func (this *BaseWebController) AdminAuth() {
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

func (this *BaseWebController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseWebController) display(tpl ...string) {
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

func (this *BaseWebController) getBgAreaAction(action string) string {
	return "area/" + action
}

func (this *BaseWebController) getBgWebAction(action string) string {
	return "web/" + action
}

func (this *BaseWebController) getBgToolAction(action string) string {
	return "tools/" + action
}

func (this *BaseWebController) getBgAppAction(action string) string {
	return this.getBgTestAction("app/" + action)
}

func (this *BaseWebController) getBgTestAction(action string) string {
	return "test/" + action
}

func (this *BaseWebController) getBgApiAction(action string) string {
	return "api/" + action
}

func (this *BaseWebController) ajaxMsg(msg interface{}, msgNo int) {
	out := make(map[string]interface{})
	out["status"] = msgNo
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseWebController) ajaxList(msg interface{}, msgNo int, count int64, data interface{}) {
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

func (this *BaseWebController) row(row map[string]interface{}, model interface{}, isEdit bool) {
	this.Data["row"] = this.relectRow(row, model, 0, isEdit)
}

func (this *BaseWebController) parse(list []map[string]interface{}, row map[string]interface{}, k int, v interface{}, isEdit bool) {
	if list == nil {
		return
	}
	list[k] = this.relectRow(row, v, 0, isEdit)
}

func (this *BaseWebController) group(list []map[string]interface{}, row map[string]interface{}, k int, v interface{}, id int64, isEdit bool) {
	if list == nil {
		return
	}
	list[k] = this.relectRow(row, v, id, isEdit)
}

func (this *BaseWebController) relectRow(row map[string]interface{}, v interface{}, id int64, isEdit bool) map[string]interface{} {
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
