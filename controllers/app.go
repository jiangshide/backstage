package controllers

import (
	"time"
	"strings"
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"os/exec"
	"github.com/jiangshide/GoComm/utils"
	"strconv"
	"github.com/Mr4Mike4/ApkInfGo"
	"github.com/skip2/go-qrcode"
	"github.com/astaxie/beego/orm"
	"backstage/models"
)

type AppController struct {
	BaseController
}

func (this *AppController) List() {
	this.pageTitle("应用列表~状态:1:打包中,2:打包成功,3:打包失败,4,无")
	channel := new(models.Channel)
	result, _ := channel.List(-1, 0)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.Data["channelList"] = list
	this.display("app/list")
}

func (this *AppController) Add() {
	this.display("app/list")
}

func (this *AppController) Table() {
	app := new(models.App)
	result, count := app.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AppController) AjaxSave() {
	channel := new(models.Channel)
	channel.Id = this.getInt64("channel", 0)
	isNet := this.getString("isNet")
	beego.Info("----------id:", channel.Id, " | isNet:", isNet)
	if err := channel.Query().One(channel); err == nil {
		project := this.getString("project")
		version := this.getString("version")
		environment := this.getString("environment")
		buildType := this.getString("buildType")
		proguard := this.getString("proguard")
		this.pack(channel.Platform, strconv.FormatBool(proguard == "on"),project, channel.Name, channel.Pkg, channel.FriendId, channel.Channel, environment, buildType, version)
	} else {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
}

func (this *AppController) AjaxDel() {
	app := new(models.App)
	app.Id = this.getId()
	if err := app.Query();err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	if _,err := app.Del();err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	this.delFile(app.AppUrl,app.QrUrl)
	this.ajaxMsg("",MSG_OK)
}

/**
/Users/glzc/developer/gradle/gradle-4.4/bin/gradle -POUTPUT_FILE=/Users/glzc/developer/server/go/src/etongdai_web/static/app/apkUrl/ -PAPPLICATION_ID=com.stateunion.p2p.etongdai -PIS_REMOTE=true -PenableProguard=true -PAPP_NAME=易通贷理财1 -PFRINED_ID=ODM1MzEyNw== -PAPP_VERSION_NAME=3.0.18 -PAPP_VERSION_CODE=3019 -PENVIRONMENT=production -PUMENG_VALUE=test1990 -PLABLE=杰克 -PDATE= -PBUILD_SCRIPT=/Users/glzc/developer/client/android/Android/app/build.gradle assembleRelease --stacktrace   -b /Users/glzc/developer/client/android/Android/app/build.gradle
 */

func (this *AppController) pack(platform int,proguard, project, name, pkg, friendId, channel, environment, buildType, version string) {
	var err error
	gradle := "/Users/glzc/developer/gradle/gradle-4.4/bin/gradle"
	sourceVersionCode := strings.Replace(version, ".", "", -1)
	versionCodeNum, _ := strconv.Atoi(sourceVersionCode)
	versionCodeNum = versionCodeNum + 1
	versionCode := strconv.Itoa(versionCodeNum)
	author := this.userName
	date := fmt.Sprint(time.Now().Unix())
	app := new(models.App)
	app.Name = name
	app.Pkg = pkg
	app.FriendId = friendId
	app.Channel = channel
	app.Version = version
	app.VersionCode = versionCodeNum
	app.Environment = environment
	app.Platform = platform
	app.Status = 1
	app.CreateId = this.userId
	app.CreateTime = time.Now().Unix()
	_, err = app.Add()
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	cmd := gradle + " -POUTPUT_FILE=" + utils.GetCurrentDir("/static/app/apkUrl") +" -PenableProguard="+proguard+ " -PAPPLICATION_ID=" + pkg + " -PIS_REMOTE=true -PAPP_NAME=" + name + " -PFRINED_ID=" + friendId + " -PAPP_VERSION_NAME=" + version + " -PAPP_VERSION_CODE=" + versionCode + " -PENVIRONMENT=" + environment + " -PUMENG_VALUE=" + channel + " -PLABLE=" + author + " -PDATE=" + date + " -PBUILD_SCRIPT=" + project + "/android/app/build.gradle" +
		" assemble" + buildType + " -b " + project + "/app/build.gradle"
	app.Cmd = cmd
	log, err := this.ExeCommand(cmd)
	app.Log = log
	if err != nil {
		app.ErrMsg = err.Error()
		app.Status = 3
		app.Update()
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	err, apkName,size := this.copyFile(project, author, channel, version, versionCode, date, environment, buildType, "/static/app/apkUrl/")
	beego.Info("-------err:", err, " | apkName:", apkName)
	if err != nil {
		app.ErrMsg = err.Error()
		app.Status = 3
		_, err = app.Update()
		this.ajaxMsg(err.Error(), MSG_ERR)
		//this.redirect(beego.URLFor("AppController.List"))
	}
	app.Size = size
	err, qrImgPath, apkPathUrl := this.generalQrImg(apkName)
	if err != nil {
		app.ErrMsg = err.Error()
		app.Status = 3
		app.Update()
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	app.Status = 2
	app.QrUrl = qrImgPath
	app.AppUrl = apkPathUrl
	app.UpdateTime = time.Now().Unix()
	if _, err := app.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AppController) copyFile(project, author, channelName, appVersion, appVersionCode, date, environment, buildType, pathSuffix string) (error, string,int64) {
	apkName := "etongdai-" + author + "-android-" + channelName + "-v" + appVersion + "-" + appVersionCode + "-" + date + "-" + environment + "-" + buildType + ".apk"
	destApk := utils.GetCurrentDir(pathSuffix) + apkName
	sourceApk := project + "/app/build/outputs/apk/release/" + apkName
	size, err := utils.CopyFile(destApk, sourceApk)
	beego.Info("---------size:",size)
	if err == nil {
		os.Remove(sourceApk)
	}
	return err, apkName,size
}

func (this *AppController) ExeCommand(command string) (string, error) {
	res, err := exec.Command("/bin/sh", "-c", ``+command+``).Output()
	return string(res), err
}

func (this *AppController) generalQrImg(apkName string) (error, string, string) {
	rootPath := utils.GetCurrentDir("")
	qrImgPath := rootPath + "/static/app/qrImgUrl/"
	if err := utils.Exist(qrImgPath); err != nil {
		if err := os.MkdirAll(qrImgPath, os.ModePerm); err != nil {
			beego.Error(err)
		}
	}
	currDate := strconv.FormatInt(time.Now().UnixNano(), 13)
	qrImgPath = qrImgPath + currDate + ".png"
	qrImgUrl := "/static/app/qrImgUrl/" + currDate + ".png"
	apkPathUrl := "/static/app/apkUrl/" + apkName
	return qrcode.WriteFile("http://"+utils.GetLocalAdder()+":"+beego.AppConfig.String("httpport")+apkPathUrl, qrcode.Medium, 256, qrImgPath), qrImgUrl, apkPathUrl
}

func (this *AppController) GetApkInfo(filePath string) ApkInfGo.ApkInfoSt {
	aapt := "/Users/glzc/Library/Android/sdk/build-tools/27.0.3/aapt"
	keytool := "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/bin/keytool"
	apk := ApkInfGo.ApkInfo(aapt).CertKeyTool(keytool).File(filePath)
	return *apk
}

type UpdateController struct {
	BaseController
}

var updateMode = []string{"普通提示更新", "提示强制更新", "后台自动下载后更新(非静默更新)", "静默更新"}

func (this *UpdateController) channel() {
	this.Data["status"] = updateMode
	channel := new(models.Channel)
	result, _ := channel.List(-1, 0)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.Data["channelList"] = list
}

func (this *UpdateController) List() {
	this.pageTitle("更新列表")
	this.display("app/update/list")
}

func (this *UpdateController) Add() {
	this.pageTitle("新增")
	this.channel()
	this.display("app/update/add")
}

func (this *UpdateController) Edit() {
	this.pageTitle("编辑渠道")
	update := new(models.Update)
	update.Id = this.getId()
	if err := update.Query().One(update); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, update)
	this.channel()
	this.Data["channelName"] = update.Channel
	this.Data["modeName"] = update.Status
	this.display("app/update/edit")
}

func (this *UpdateController) AjaxSave() {
	update := new(models.Update)
	update.Id = this.getId()
	update.Name = this.getString("name")
	update.Url = this.getString("url")
	beego.Info("----------url:", update.Url)
	channel := new(models.Channel)
	channel.Id = this.getInt64("channelId", 0)
	beego.Info("-----------------id:", channel.Id)
	errs := channel.Query().Filter("id",channel.Id).One(channel)
	beego.Info("----------channel:",channel.Channel)
	if errs == nil {
		update.Channel = channel.Channel
		update.Platform = channel.Platform
	}
	update.Version = this.getString("version")
	update.Status = this.getInt("status", 0)
	update.Content = this.getString("content")
	var err error
	if update.Id == 0 {
		update.CreateId = this.userId
		update.CreateTime = time.Now().Unix()
		_, err = update.Add()
	} else {
		update.CreateId = this.getInt64("createId", 0)
		update.CreateTime = this.getInt64("createTime", 0)
		update.UpdateId = this.userId
		update.UpdateTime = time.Now().Unix()
		_, err = update.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *UpdateController) Table() {
	update := new(models.Update)
	result, count := update.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.Data["updateMode"] = updateMode
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *UpdateController) AjaxDel() {
	update := new(models.Update)
	update.Id = this.getId()
	if err := update.Query().One(update);err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	if _, err := update.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.delFile(update.Url)
	this.ajaxMsg("", MSG_OK)
}

type StopController struct {
	BaseController
}

func (this *StopController) List() {
	this.pageTitle("停服列表")
	this.display("app/stop/list")
}

func (this *StopController) Add() {
	this.pageTitle("新增")
	this.Data["platforms"] = platforms
	this.groupBy()
	this.display("app/stop/add")
}

func (this *StopController) groupBy() {
	var apps [] models.Channel
	orm.NewOrm().Raw("select name from zd_app_channel group by name").QueryRows(&apps)
	this.Data["apps"] = apps
}

func (this *StopController) Edit() {
	this.pageTitle("编辑停服内容")
	stop := new(models.Stop)
	stop.Id = this.getId()
	beego.Info("-------------id:", stop.Id)
	if err := stop.Query().Filter("id", stop.Id).One(stop); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, stop)
	this.Data["platform"] = stop.Platform
	this.Data["platforms"] = platforms
	this.groupBy()
	this.Data["app"] = stop.Name
	this.display("app/stop/edit")
}

func (this *StopController) AjaxSave() {
	stop := new(models.Stop)
	stop.Id = this.getId()
	beego.Info("-------------id:", stop.Id)
	stop.Name = this.getString("name")
	stop.Channel = this.getString("channel")
	stop.Version = this.getString("version")
	stop.Url = this.getString("url")
	stop.Platform = this.getInt("platform", 0)
	stop.Url = this.getString("url")
	var err error
	if stop.Id == 0 {
		stop.CreateId = this.userId
		stop.CreateTime = time.Now().Unix()
		_, err = stop.Add()
	} else {
		stop.CreateId = this.getInt64("createId", 0)
		stop.CreateTime = this.getInt64("createTime", 0)
		stop.UpdateId = this.userId
		stop.UpdateTime = time.Now().Unix()
		_, err = stop.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *StopController) Table() {
	stop := new(models.Stop)
	result, count := stop.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *StopController) AjaxDel() {
	stop := new(models.Stop)
	stop.Id = this.getId()
	if err := stop.Query().One(&stop);err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	if _, err := stop.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.delFile(stop.Url)
	this.ajaxMsg("", MSG_OK)
}

type ChannelController struct {
	BaseController
}

var ChannelName = map[string]string{"yingxiao1037": "MTA3MTQwNTY=", "oppo": "MjExOTk5NQ==", "vivo": "MjExOTIwMQ==", "vivo马甲包": "OTY0NDgwNA==", "yingxiao1031": "MTA3MTQwNTA=", "test1990": "ODM1MzEyNw==", "yingxiao1016": "MTA3MTQwMzU=", "yingxiao1012": "MTA3MTQwMzE=", "yingxiao1001": "MTA3MTQwMjA=", "yingxiao1002": "MTA3MTQwMjE=", "xsh60": "OTY0NDY2OA==", "shangwu80": "MTAyNjkyNjI=", "shangwu32": "MTAyNjgzNzQ=", "华为": "MjExOTEzOQ==", "2345手机助手": "MjExOTU3Mw==", "机锋网": "MjExOTU5Mw==",
	"厉趣网": "MjExOTYzMQ==", "360应用市场": "MjExOTk2NQ==", "豌豆荚": "MjEyMDAyNQ==",
	"乐商店": "MjEyMDA0NQ==", "小沃科技（联通）": "MjExOTQ3NQ==", "小米应用市场": "MjEyMDA2NQ==", "乐视": "MjEyMDA4Nw==",
	"木蚂蚁": "MjExOTMwOQ==", "优亿": "MjEyMDEyNQ==", "应用宝应用市场": "MjEyMDIyNQ==", "魅族": "MjEyMDI4MQ==",
	"金立": "MjExOTI1NQ==", "应用汇": "MjEyMDMxNw==", "搜狗": "MjEyMDUwMQ==", "应用宝CPD": "OTg0NTQ4Ng==",
	"三星": "OTY0NDY4MA==", "百度": "OTY0NDcwMA==", "360应用市场马甲包": "OTY0NDgxMg=="}
var statuss = map[int]string{1: "正常", 2: "下架", 3: "异常"}
var platforms = []string{"无限制", "android", "ios", "web"}

func (this *ChannelController) InitChannelData() {
	for k, v := range ChannelName {
		channel := new(models.Channel)
		channel.Name = "易通贷理财"
		channel.Pkg = "com.stateunion.p2p.etongdai"
		channel.FriendId = v
		channel.Channel = k
		channel.Platform = 1
		channel.Status = 1
		channel.CreateId = this.userId
		channel.UpdateId = channel.CreateId
		channel.CreateTime = time.Now().Unix()
		channel.UpdateTime = channel.CreateTime
		if k == "oppo" {
			channel.Name = "易通贷理财-投资理财"
		}
		if strings.Contains(k, "马甲包") {
			channel.Name = "易理财"
			channel.Pkg = "com.stateunion.p2p.etongdai.vest"
		}
		channel.Add();
	}
}

func (this *ChannelController) List() {
	this.pageTitle("渠道列表")
	//this.InitChannelData()
	this.display("app/channel/list")
}

func (this *ChannelController) Add() {
	this.pageTitle("新增")
	this.Data["statuss"] = statuss
	this.Data["platforms"] = platforms
	this.display("app/channel/add")
}

func (this *ChannelController) Edit() {
	this.pageTitle("编辑渠道")
	channel := new(models.Channel)
	channel.Id = this.getId()
	if err := channel.Query().One(&channel); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, channel)
	this.Data["platform"] = channel.Platform
	this.Data["status"] = channel.Status
	this.Data["statuss"] = statuss
	this.Data["platforms"] = platforms
	this.display("app/channel/edit")
}

func (this *ChannelController) AjaxSave() {
	channel := new(models.Channel)
	channel.Id = this.getId()
	channel.Name = this.getString("name")
	channel.Pkg = this.getString("pkg")
	channel.FriendId = this.getString("friendId")
	channel.Channel = this.getString("channel")
	channel.Platform = this.getInt("platform", 0)
	channel.Status = this.getInt("status", 1)
	var err error
	if channel.Id == 0 {
		channel.CreateId = this.userId
		channel.CreateTime = time.Now().Unix()
		_, err = channel.Add()
	} else {
		channel.CreateId = this.getInt64("createId", 0)
		channel.CreateTime = this.getInt64("createTime", 0)
		channel.UpdateId = this.userId
		channel.UpdateTime = time.Now().Unix()
		_, err = channel.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ChannelController) Table() {
	channel := new(models.Channel)
	result, count := channel.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ChannelController) AjaxDel() {
	channel := new(models.Channel)
	channel.Id = this.getId()
	if _, err := channel.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

var advertTypes = map[int]string{1: "开机广告", 2: "app启动广告", 3: "app第一安装导航广告", 4: "其它"}

type AdvertController struct {
	BaseController
}

func (this *AdvertController) List() {
	this.pageTitle("广告列表")
	this.display("app/advert/list")
}

func (this *AdvertController) Add() {
	this.pageTitle("新增")
	this.Data["advertTypes"] = advertTypes
	this.Data["statuss"] = statuss
	this.Data["platforms"] = platforms
	this.display("app/advert/add")
}

func (this *AdvertController) Edit() {
	this.pageTitle("编辑广告")
	advert := new(models.Advert)
	advert.Id = this.getId()
	if err := advert.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, advert)
	this.Data["advertTypes"] = advertTypes
	this.Data["advertType"] = advert.Type
	this.Data["platform"] = advert.Platform
	this.Data["platforms"] = platforms
	this.Data["statuss"] = statuss
	this.Data["status"] = advert.Status
	this.display("app/advert/edit")
}

func (this *AdvertController) AjaxSave() {
	beego.Info("---------------ajaxSave")
	advert := new(models.Advert)
	advert.Id = this.getId()
	advert.Name = this.getString("name")
	advert.Type = this.getInt("type", 0)
	advert.Platform = this.getInt("platform", 0)
	advert.Price = this.getInt("price", 0)
	advert.Times = this.getInt("times", 0)
	advert.Url = this.getString("file")
	advert.Status = this.getInt("status", 1)
	var err error
	if advert.Id == 0 {
		advert.CreateId = this.userId
		advert.CreateTime = time.Now().Unix()
		_, err = advert.Add()
	} else {
		advert.CreateId = this.getInt64("createId", 0)
		advert.CreateTime = this.getInt64("createTime", 0)
		advert.UpdateId = this.userId
		advert.UpdateTime = time.Now().Unix()
		_, err = advert.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AdvertController) Table() {
	advert := new(models.Advert)
	result, count := advert.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AdvertController) AjaxDel() {
	advert := new(models.Advert)
	advert.Id = this.getId()
	if err := advert.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR);
	}
	if _, err := advert.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.delFile(advert.Url)
	this.ajaxMsg("", MSG_OK)
}
