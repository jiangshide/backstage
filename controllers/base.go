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
	"net"
	"github.com/pkg/errors"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
	AUTH    = "auth"
)

type BaseController struct {
	beego.Controller
	version    string
	controller string
	action     string
	userId     int64
	userName   string
	parentId   int64
	logo       string
	page       int
	pageSize   int
	offSet     int
	allowUrl   string
	upload     string
}

var defaultTips = "该项不能为空!"
var defaultMinSize = 3

func (this *BaseController) Prepare() {
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	beego.Info("-----------controller:", this.controller, " | action:", this.action)
	if this.controller == "permissionuser" && this.action == "login" {
		if this.userId > 0 {
			this.redirect("/")
		}
	} else {
		this.Data["siteName"] = beego.AppConfig.String("site.app_name")
		this.version = beego.AppConfig.String("version")
		this.page, _ = this.GetInt("page", 1)
		this.pageSize, _ = this.GetInt("limit", 30)
		this.offSet = (this.page - 1) * this.pageSize
		Tinify.SetKey(this.GetString("pic_key"))
		this.upload = this.GetString("upload", "/static/upload/")

		if user, err := this.auth(); user != nil {
			this.userId = user.Id
			this.userName = user.Name
			this.Data["userId"] = this.userId
			this.Data["userName"] = this.userName
			var parentIds int64
			parentIds, _ = beego.AppConfig.Int64("parentId")
			if parentIds > 0 {
				this.parentId = parentIds
			}

			beego.Info("--------userName:", this.userName, " | userId:", this.userId, " | upload:", this.upload, " | parentId:", this.parentId)
		} else {
			beego.Error(err)
			this.redirect(beego.URLFor("UserController.Login"))
		}
	}

}

func (this *BaseController) parentIds() {
	parentId := this.getInt64("parentId", 0)
	if parentId > 0 {
		beego.AppConfig.Set("parentId", fmt.Sprint(parentId))
	}
}

func (this *BaseController) setCook(user *models.PermissionUser, time int) {
	this.Ctx.SetCookie(AUTH, fmt.Sprint(user.Id)+"|"+utils.Md5(this.getIp()+"|"+user.Password+user.Salt), time)
}

func (this *BaseController) auth() (users *models.PermissionUser, errs error) {
	cook := this.Ctx.GetCookie(AUTH)
	if strings.Contains(cook, "|") {
		cookArr := strings.Split(cook, "|")
		user := new(models.PermissionUser)
		user.Id, _ = strconv.ParseInt(cookArr[0], 11, 64)
		if errs = user.Query(); errs == nil {
			users = user
			auth := new(models.PermissionAuth)
			if result, count := auth.List(-1, 0); count > 0 {
				sideMenu := make([]map[string]interface{}, count)
				sideMenuChild := make([]map[string]interface{}, count)
				i, j := 0, 0
				for _, v := range result {
					menu := make(map[string]interface{}, 0)
					menu["Id"] = v.Id
					menu["Pid"] = v.Pid
					menu["Name"] = v.Name
					menu["Action"] = v.Action
					menu["Icon"] = v.Icon
					menu["IsShow"] = v.IsShow
					if v.IsShow == 1 {
						if v.Pid == 1 {
							sideMenu[i] = menu
							i++
						} else {
							sideMenuChild[j] = menu
							j++
						}
					}
				}
				this.Data["sideMenu"] = sideMenu[:i]
				this.Data["sideMenuChild"] = sideMenuChild[:j]
			} else {
				errs = errors.New("还没有数据!")
			}
		}
	}
	return
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, err := this.GetInt(key, defaultValue)
	if err != nil {
		beego.Info("------------key:", key, " | err:", err)
		this.ajaxMsg(err, MSG_ERR)
	}
	return resInt
}

func (this *BaseController) getInt64(key string, defaultValue int64) int64 {
	resInt, err := this.GetInt64(key, defaultValue)
	if err != nil {
		beego.Info("--------key:", key, " | err:", err)
		this.ajaxMsg(err, MSG_ERR)
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

func (this *BaseController) getString(key, tips string, minSize int) (value string) {
	value = strings.TrimSpace(this.GetString(key, ""))
	errorMsg := ""
	if len(value) == 0 {
		errorMsg = tips
	} else if len(value) < minSize {
		errorMsg = "长度不能小于" + fmt.Sprint(minSize) + "字符:" + value
	}
	if errorMsg != "" {
		beego.Info("--------------key:", key, " | errorMsg:", errorMsg)
		this.ajaxMsg(errorMsg, MSG_ERR)
	}
	return
}

func (this *BaseController) pageTitle(pageTitle string) {
	this.Data["pageTitle"] = pageTitle
}

func (this *BaseController) redirect(url string) {
	beego.Info("------------redirect:", url)
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) display(tpl ...string) {
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = this.controller + "/" + this.action + ".html"
	}
	this.Layout = "comm/layout.html"
	this.TplName = tplName
}

func (this *BaseController) getBgPermissionAction(action string) string {
	return "permission/" + action
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
	beego.Info("-----------1----msg:",msg," | msgNo",msgNo," | count:",count," | data:",data)
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) row(row map[string]interface{}, model interface{}, isEdit bool, displayUrl string) {
	this.Data["row"] = this.relectRow(row, model, 0, isEdit)
	beego.Info("displayUrl:", displayUrl)
	this.display(displayUrl)
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

	var parentId int64
	for i := 0; i < size; i++ {
		v := value.Field(i)
		fieldName = field.Field(i).Name
		link := field.Field(i).Tag.Get("link")
		fieldName = utils.StrFirstToLower(fieldName)
		if fieldName == "id" {
			if id > 0 {
				if v.Int() == id {
					row["selected"] = true
				} else {
					row["selected"] = false
				}
			}
			parentId = v.Int()
			//continue
		}

		switch value.Field(i).Kind() {
		case reflect.Bool:
			row[fieldName] = v.Bool()
		case reflect.Int:
			if v.Int() > 0 {
				row[fieldName] = v.Int()
			}
		case reflect.Int64:
			if fieldName == "createTime" {
				if v.Int() > 0 && isEdit {
					row[fieldName] = v.Int()
				} else {
					row["createTimeFormat"] = beego.Date(time.Unix(v.Int(), 0), "Y-m-d H:i:s")
				}
				continue
			}
			if fieldName == "updateTime" {
				if v.Int() > 0 && !isEdit {
					row["updateTimeFormat"] = beego.Date(time.Unix(v.Int(), 0), "Y-m-d H:i:s")
				}
				continue
			}
			if fieldName == "createId" || fieldName == "updateId" {
				if v.Int() > 0 {
					if isEdit {
						if fieldName == "createId" {
							row[fieldName] = v.Int()
						}
					} else {
						admin := new(models.PermissionUser)
						admin.Id = v.Int()
						if err := admin.Query(); err == nil {
							if fieldName == "createId" {
								row["createName"] = admin.Name
							} else {
								row["updateName"] = admin.Name
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
				if strings.Contains(v.String(), ".png") || strings.Contains(v.String(), ".jpg") {
					row["file"] = "<img src='" + v.String() + "' />"
				} else {
					row["file"] = v.String()
				}
				this.setFileSize(row, v.String())
			} else {
				if len(link) > 0 && !isEdit {
					row[fieldName] = "<a href='" + link + v.String() + "?parentId=" + fmt.Sprint(parentId) + "' style='color:#6495ED;'>" + v.String() + "</a>"
				} else {
					row[fieldName] = v.String()
				}
			}
		}
	}

	beego.Info("-------row:", row, " | isEdit:", isEdit)
	return row
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getIp() string {
	return this.Ctx.Input.IP()
}

func (this *BaseController) getMac() (mac string) {
	inter, err := net.InterfaceByName("eth0")
	if err != nil {
		beego.Info("err:", err)
	} else {
		mac = inter.HardwareAddr.String()
	}
	return
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

func (this *BaseController) showTips(errorMsg interface{}) {
	flash := beego.NewFlash()
	flash.Error(fmt.Sprint(errorMsg))
	flash.Store(&this.Controller)
	controller, action := this.GetControllerAndAction()
	this.Redirect(beego.URLFor(controller+"."+action), 302)
	this.StopRun()
}
