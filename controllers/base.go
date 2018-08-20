package controllers

import (
	"strconv"
	"strings"
	"github.com/jiangshide/GoComm/utils"
	"github.com/astaxie/beego"
	"backstage/models"
	"reflect"
	"time"
	"fmt"
	"github.com/jiangshide/tinify-go/tinify"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.User
	userId         int64
	userName       string
	loginName      string
	page           int
	pageSize       int
	offSet         int
	allowUrl       string
	upload         string
	project        string
}

//前期准备
func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	beego.Info("----------------controllerName:", controllerName, " | actionName:", actionName)
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.page, _ = this.GetInt("page", 1)
	this.pageSize, _ = this.GetInt("limit", 20)
	this.offSet = (this.page - 1) * this.pageSize
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.auth()
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
	this.upload = beego.AppConfig.String("upload")
	this.project = beego.AppConfig.String("project")
}

func (this *BaseController) pageTitle(pageTitle string) {
	this.Data["pageTitle"] = pageTitle
}

//登录权限验证
func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.ParseInt(idstr, 10, 64)
		if userId > 0 {
			admin := new(models.User)
			user, err := admin.GetById(userId)
			if err == nil && password == utils.Md5(this.getClientIp()+"|"+user.Password+user.Salt) {
				this.userId = user.Id
				this.loginName = user.Name
				this.user = user
				this.AdminAuth()
			}
			//isHasAuth := strings.Contains(this.allowUrl, this.controllerName+"/"+this.actionName)
			//noAuth := "ajaxsave/ajaxdel/table/login/quit/getnodes/start/show/ajaxapisave/index/group/"
			//isNoAuth := strings.Contains(noAuth, this.actionName)
			//if isHasAuth == false && isNoAuth == false {
			//	this.Ctx.WriteString("没有权限")
			//	this.ajaxMsg("没有权限", MSG_ERR)
			//	return
			//}
		}
	}

	if this.userId == 0 && this.actionName == "index" {
		this.redirect(beego.URLFor("UserController.Login"))
	}
}

func (this *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if this.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	sideMenu := make([]map[string]interface{}, len(result))
	sideMenuChild := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/" {
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		row["Id"] = int(v.Id)
		row["Sort"] = v.Sort
		row["AuthName"] = v.AuthName
		row["AuthUrl"] = v.AuthUrl
		row["Icon"] = v.Icon
		row["Pid"] = int(v.Pid)
		if (v.IsShow == 1) {
			if (v.Pid == 1) {
				sideMenu[i] = row
				i++
			} else {
				sideMenuChild[j] = row
				j++
			}
		}
	}

	this.Data["sideMenu"] = sideMenu[:i]
	this.Data["sideMenuChild"] = sideMenuChild[:j]

	this.allowUrl = allow_url + "/home/index"
}

func (this *BaseController) row(row map[string]interface{}, model interface{}) {
	var field reflect.Type
	var value reflect.Value
	field = reflect.TypeOf(model).Elem()
	value = reflect.ValueOf(model).Elem()
	size := field.NumField()
	var fieldName string
	if row == nil {
		row = make(map[string]interface{})
	}
	for i := 0; i < size; i++ {
		v := value.Field(i)
		fieldName = field.Field(i).Name
		this.Data[fieldName] = v
	}
}

func (this *BaseController) parse(list []map[string]interface{}, row map[string]interface{}, k int, v interface{}, isEdit bool) {
	if list == nil {
		return
	}
	list[k] = this.relectRow(row, v, 0, isEdit)
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
						user := new(models.User)
						user.Id = v.Int()
						if err := user.Query(); err == nil {
							if fieldName == "createId" {
								row["createName"] = user.Name
							} else {
								row["updateName"] = user.Name
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
					beego.Info("----------------img:", v.String())
					row["img"] = "<img src='" + v.String() + "' />"
				} else {
					row["file"] = "<a href='" + "http://" + utils.GetLocalAdder() + ":" + beego.AppConfig.String("httpport") + v.String() + "' >Download</>"
				}
				row[fieldName] = v.String()
				this.setFileSize(row, utils.GetCurrentDir(v.String()))
			} else {
				if len(link) > 0 && !isEdit {
					row[fieldName] = "<a href='" + link + v.String() + "?parentId=" + fmt.Sprint(parentId) + "' style='color:#6495ED;'>" + v.String() + "</a>"
				} else {
					row[fieldName] = v.String()
				}
			}
		}
	}
	//beego.Info("-------row:", row, " | isEdit:", isEdit)
	return row
}

func (this *BaseController) getId() int64 {
	id, err := this.GetInt64("id", 0)
	if err != nil {
		this.ajaxMsg(err, MSG_ERR)
	}
	return id
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, _ := this.GetInt(key, defaultValue)
	return resInt
}

func (this *BaseController) getInt64(key string, defaultValue int64) int64 {
	resInt64, _ := this.GetInt64(key, defaultValue)
	return resInt64
}

func (this *BaseController) getString(key string) string {
	return strings.TrimSpace(this.GetString(key, ""))
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := this.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	beego.Info("------------url:", url)
	this.StopRun()
}

//加载模板
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.Layout = "public/layout.html"
	this.TplName = tplname
}

//ajax返回
func (this *BaseController) ajaxMsgTips(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
}

//ajax返回
func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//ajax返回 列表
func (this *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
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

//分组公共方法
type groupList struct {
	Id        int
	GroupName string
}

//获取单个分组信息
func getGroupInfo(gl []groupList, groupId int) (groupInfo groupList) {
	for _, v := range gl {
		if v.Id == groupId {
			groupInfo = v
		}
	}
	return
}

type sourceList struct {
	Id         int
	SourceName string
	GroupId    int
	GroupName  string
}
