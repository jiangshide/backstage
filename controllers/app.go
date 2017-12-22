package controllers

import (
	"zd112_backstage/models"
	"time"
)

type ChannelController struct {
	BaseWebController
}

func (this *ChannelController) List() {
	this.pageTitle("渠道列表")
	this.display(this.getBgAppAction("channel/list"))
}

func (this *ChannelController) Add() {
	this.pageTitle("增加渠道名称")
	this.display(this.getBgAppAction("channel/add"))
}

func (this *ChannelController) Edit() {
	this.pageTitle("编辑渠道名称")
	channel := new(models.Channel)
	channel.Id = this.getId64(0)
	if err := channel.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, channel, true)
	this.display(this.getBgAppAction("channel/edit"))
}

func (this *ChannelController) AjaxSave() {
	channel := new(models.Channel)
	channel.Id = this.getId64(0)
	channel.Name = this.getString("name", "名称不能为空!", 1)
	channel.FriendId = this.getString("friend_id", "", 0)
	channel.Descript = this.getString("descript", "", 0)
	var err error
	if channel.Id == 0 {
		channel.CreateId = this.userId
		channel.CreateTime = time.Now().Unix()
		_, err = channel.Add()
	} else {
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
	channel.Id = this.getId64(0)
	if _, err := channel.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type AppNameController struct {
	BaseWebController
}

func (this AppNameController) List() {
	this.pageTitle("应用名称列表")
	this.display(this.getBgAppAction("app_name/list"))
}

func (this AppNameController) Add() {
	this.pageTitle("增加应用名称")
	this.display(this.getBgAppAction("app_name/add"))
}

func (this AppNameController) Edit() {
	this.pageTitle("编辑应用名称")
	appName := new(models.AppName)
	appName.Id = this.getId64(0)
	if err := appName.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, appName, true)
	this.display(this.getBgAppAction("app_name/edit"))
}

func (this AppNameController) AjaxSave() {
	appName := new(models.AppName)
	appName.Id = this.getId64(0)
	appName.Name = this.getString("name", "应用名称不能为空!", 1)
	appName.Descript = this.getString("descript", "", 0)
	var err error
	if appName.Id == 0 {
		appName.CreateId = this.userId
		appName.CreateTime = time.Now().Unix()
		_, err = appName.Add()
	} else {
		appName.UpdateId = this.userId
		appName.UpdateTime = time.Now().Unix()
		_, err = appName.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this AppNameController) Table() {
	appName := new(models.AppName)
	result, count := appName.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AppNameController) AjaxDel() {
	appName := new(models.AppName)
	appName.Id = this.getId64(0)
	if _, err := appName.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type PkgsController struct {
	BaseWebController
}

func (this PkgsController) List() {
	this.pageTitle("应用包列表")
	this.display(this.getBgAppAction("pkgs/list"))
}

func (this PkgsController) Add() {
	this.pageTitle("增加应用包")
	this.display(this.getBgAppAction("pkgs/add"))
}

func (this PkgsController) Edit() {
	this.pageTitle("编辑应用包")
	pkgs := new(models.Pkgs)
	pkgs.Id = this.getId64(0)
	if err := pkgs.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, pkgs, true)
	this.display(this.getBgAppAction("pkgs/edit"))
}

func (this *PkgsController) AjaxSave() {
	pkg := new(models.Pkgs)
	pkg.Id = this.getId64(0)
	pkg.Name = this.getString("name", "包名称不能为空!", 1)
	pkg.Descript = this.getString("descript", "", 0)
	var err error
	if pkg.Id == 0 {
		pkg.CreateId = this.userId
		pkg.CreateTime = time.Now().Unix()
		_, err = pkg.Add()
	} else {
		pkg.UpdateId = this.userId
		pkg.UpdateTime = time.Now().Unix()
		_, err = pkg.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *PkgsController) Table() {
	pkg := new(models.Pkgs)
	result, count := pkg.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *PkgsController) AjaxDel() {
	pkg := new(models.Pkgs)
	pkg.Id = this.getId64(0)
	if _, err := pkg.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type VersionController struct {
	BaseWebController
}

func (this *VersionController) List() {
	this.pageTitle("应用版本列表")
	this.display(this.getBgAppAction("version/list"))
}

func (this *VersionController) Add() {
	this.pageTitle("增加应用版本")
	this.display(this.getBgAppAction("version/add"))
}

func (this *VersionController) Edit() {
	this.pageTitle("编辑应用版本")
	version := new(models.Version)
	version.Id = this.getId64(0)
	if err := version.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, version, true)
	this.display(this.getBgAppAction("version/edit"))
}

func (this *VersionController) AjaxSave() {
	version := new(models.Version)
	version.Id = this.getId64(0)
	version.Name = this.getString("name", "版本名称不能为空!", 1)
	version.Descript = this.getString("descript", "", 0)
	var err error
	if version.Id == 0 {
		version.CreateId = this.userId
		version.CreateTime = time.Now().Unix()
		_, err = version.Add()
	} else {
		version.UpdateId = this.userId
		version.UpdateTime = time.Now().Unix()
		_, err = version.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *VersionController) Table() {
	version := new(models.Version)
	result, count := version.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *VersionController) AjaxDel() {
	version := new(models.Version)
	version.Id = this.getId64(0)
	if _, err := version.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CodeController struct {
	BaseWebController
}

func (this *CodeController) List() {
	this.pageTitle("应用编号列表")
	this.display(this.getBgAppAction("code/list"))
}

func (this *CodeController) Add() {
	this.pageTitle("增加应用编号列表")
	this.display(this.getBgAppAction("code/add"))
}

func (this *CodeController) Edit() {
	this.pageTitle("编辑应用编号")
	code := new(models.Code)
	code.Id = this.getId64(0)
	if err := code.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, code, true)
	this.display(this.getBgAppAction("code/edit"))
}

func (this *CodeController) AjaxSave() {
	code := new(models.Code)
	code.Id = this.getId64(0)
	code.Code = this.getInt("code", 0)
	code.Descript = this.getString("descript", "", 0)
	var err error
	if code.Id == 0 {
		code.CreateId = this.userId
		code.CreateTime = time.Now().Unix()
		_, err = code.Add()
	} else {
		code.UpdateId = this.userId
		code.UpdateTime = time.Now().Unix()
		_, err = code.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CodeController) Table() {
	code := new(models.Code)
	result, count := code.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CodeController) AjaxDel() {
	code := new(models.Code)
	code.Id = this.getId64(0)
	if _, err := code.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type EnvController struct {
	BaseWebController
}

func (this *EnvController) List() {
	this.pageTitle("应用环境列表")
	this.display(this.getBgAppAction("env/list"))
}

func (this *EnvController) Add() {
	this.pageTitle("增加应用环境")
	this.display(this.getBgAppAction("env/add"))
}

func (this *EnvController) Edit() {
	this.pageTitle("编辑应用环境")
	env := new(models.Env)
	env.Id = this.getId64(0)
	if err := env.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, env, true)
	this.display(this.getBgAppAction("env/edit"))
}

func (this *EnvController) AjaxSave() {
	env := new(models.Env)
	env.Id = this.getId64(0)
	env.Name = this.getString("name", "应用环境名称不能为空！", 1)
	env.Descript = this.getString("descript", "", 0)
	var err error
	if env.Id == 0 {
		env.CreateId = this.userId
		env.CreateTime = time.Now().Unix()
		_, err = env.Add()
	} else {
		env.UpdateId = this.userId
		env.UpdateTime = time.Now().Unix()
		_, err = env.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *EnvController) Table() {
	env := new(models.Env)
	result, count := env.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *EnvController) AjaxDel() {
	env := new(models.Env)
	env.Id = this.getId64(0)
	if _, err := env.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type BuildController struct {
	BaseWebController
}

func (this *BuildController) List() {
	this.pageTitle("构建类型列表")
	this.display(this.getBgAppAction("build/list"))
}

func (this *BuildController) Add() {
	this.pageTitle("增加构建类型")
	this.display(this.getBgAppAction("build/add"))
}

func (this *BuildController) Edit() {
	this.pageTitle("编辑构建类型")
	build := new(models.Build)
	build.Id = this.getId64(0)
	if err := build.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, build, true)
	this.display(this.getBgAppAction("build/edit"))
}

func (this *BuildController) AjaxSave() {
	build := new(models.Build)
	build.Id = this.getId64(0)
	build.Name = this.getString("name", "名称不能为空!", 1)
	build.Descript = this.getString("descript", "", 0)
	var err error
	if build.Id == 0 {
		build.CreateId = this.userId
		build.CreateTime = time.Now().Unix()
		_, err = build.Add()
	} else {
		build.UpdateId = this.userId
		build.UpdateTime = time.Now().Unix()
		_, err = build.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *BuildController) Table() {
	build := new(models.Build)
	result, count := build.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *BuildController) AjaxDel() {
	build := new(models.Build)
	build.Id = this.getId64(0)
	if _, err := build.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TypeController struct {
	BaseWebController
}

func (this *TypeController) List() {
	this.pageTitle("构建平台列表")
	this.display(this.getBgAppAction("type/list"))
}

func (this *TypeController) Add() {
	this.pageTitle("增加平台名称")
	this.display(this.getBgAppAction("type/add"))
}

func (this *TypeController) Edit() {
	this.pageTitle("编辑平台名称")
	types := new(models.Type)
	types.Id = this.getId64(0)
	if err := types.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, types, true)
	this.display(this.getBgAppAction("type/edit"))
}

func (this *TypeController) AjaxSave() {
	types := new(models.Type)
	types.Id = this.getId64(0)
	types.Name = this.getString("name", "名称不能为空!", 1)
	types.Descript = this.getString("descript", "", 0)
	var err error
	if types.Id == 0 {
		types.CreateId = this.userId
		types.CreateTime = time.Now().Unix()
		_, err = types.Add()
	} else {
		types.UpdateId = this.userId
		types.UpdateTime = time.Now().Unix()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TypeController) Table() {
	types := new(models.Type)
	result, count := types.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TypeController) AjaxDel() {
	types := new(models.Type)
	types.Id = this.getId64(0)
	if _, err := types.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type AppController struct {
	BaseWebController
}

func (this *AppController) List() {
	this.pageTitle("应用列表")
	this.display(this.getBgTestAction("app/list"))
}

func (this *AppController) Add() {
	this.pageTitle("增加应用")
	this.setChannel(0)
	this.setBuild(0)
	this.setEnv(0)
	this.setCode(0)
	this.setVersion(0)
	this.setPkg(0)
	this.setApplication(0)
	this.setType(0)
	this.display(this.getBgTestAction("app/add"))
}

func (this *AppController) Edit() {
	app := new(models.App)
	app.Id = this.getId64(0)
	if err := app.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.setType(app.TypeId)
	this.setApplication(app.ApplicationId)
	this.setPkg(app.PkgId)
	this.setVersion(app.VersionId)
	this.setCode(app.CodeId)
	this.setEnv(app.EnvId)
	this.setBuild(app.BuildId)
	this.setChannel(app.ChannelId)
	this.row(row, app, true)
	this.display(this.getBgTestAction("app/edit"))
}

func (this *AppController) AjaxSave() {
	app := new(models.App)
	app.Id = this.getId64(0)
	app.TestId = this.getInt64("test_id", 0)
	app.Icon = this.getString("icon", "Logo不能为空!", 1)
	app.TypeId = this.getInt64("type_id", 0)
	app.ApplicationId = this.getInt64("application_id", 0)
	app.PkgId = this.getInt64("pkg_id", 0)
	app.VersionId = this.getInt64("version_id", 0)
	app.CodeId = this.getInt64("code", 0)
	app.EnvId = this.getInt64("env_id", 0)
	app.BuildId = this.getInt64("build_id", 0)
	app.ChannelId = this.getInt64("channel_id", 0)
	app.Descript = this.getString("descript", "", 0)
	app.Status = this.getInt("status", 0)
	app.Times = this.getInt("times", 0)
	app.Url = this.getString("url", "", 0)
	app.Downs = this.getInt("downs", 0)
	var err error
	if app.Id == 0 {
		app.CreateId = this.userId
		app.CreateTime = time.Now().Unix()
		_, err = app.Add()
	} else {
		app.UpdateId = this.userId
		app.UpdateTime = time.Now().Unix()
		_, err = app.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AppController) Table() {
	app := new(models.App)
	result, count := app.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		types := new(models.Type)
		types.Id = v.TypeId
		if err := types.Query(); err == nil {
			row["Type"] = types.Name
		}
		application := new(models.AppName)
		application.Id = v.ApplicationId
		if err := application.Query(); err == nil {
			row["AppName"] = application.Name
		}
		pkg := new(models.Pkgs)
		pkg.Id = v.PkgId
		if err := pkg.Query(); err == nil {
			row["Pkgs"] = pkg.Name
		}
		version := new(models.Version)
		version.Id = v.VersionId
		if err := version.Query(); err == nil {
			row["Version"] = version.Name
		}
		code := new(models.Code)
		code.Id = v.CodeId
		if err := code.Query(); err == nil {
			row["Code"] = code.Code
		}
		env := new(models.Env)
		env.Id = v.EnvId
		if err := env.Query(); err == nil {
			row["Env"] = env.Name
		}
		build := new(models.Build)
		build.Id = v.BuildId
		if err := build.Query(); err == nil {
			row["Build"] = build.Name
		}
		channel := new(models.Channel)
		channel.Id = v.ChannelId
		if err := channel.Query(); err == nil {
			row["Channel"] = channel.Name
		}
		this.parse(list, row, k, v, false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AppController) setBuild(id int64) {
	build := new(models.Build)
	result, count := build.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["BuildGroup"] = list
}

func (this *AppController) setEnv(id int64) {
	env := new(models.Env)
	result, count := env.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["EnvGroup"] = list
}

func (this *AppController) setCode(id int64) {
	code := new(models.Code)
	result, count := code.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["CodeGroup"] = list
}

func (this *AppController) setVersion(id int64) {
	version := new(models.Version)
	result, count := version.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["VersionGroup"] = list
}

func (this *AppController) setPkg(id int64) {
	pkg := new(models.Pkgs)
	result, count := pkg.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["PkgsGroup"] = list
}

func (this *AppController) setApplication(id int64) {
	application := new(models.AppName)
	result, count := application.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["AppNameGroup"] = list
}

func (this *AppController) setType(id int64) {
	types := new(models.Type)
	result, count := types.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["TypeGroup"] = list
}

func (this *AppController) setChannel(id int64) {
	channel := new(models.Channel)
	result, count := channel.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id, false)
	}
	this.Data["ChannelGroup"] = list
}

func (this *AppController) AjaxDel() {
	app := new(models.App)
	app.Id = this.getId64(0)
	if _, err := app.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
