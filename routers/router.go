package routers

import (
	"zd112_backstage/controllers"
	"github.com/astaxie/beego"
	"zd112_backstage/utils"
)

func init() {
	//beego.Router("/", &controllers.ApiDocController{}, "*:Index")
	//beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	//beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	//beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/upload", &controllers.UploadController{}, "*:Upload")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.Router("/app", &controllers.AppController{}, "*:List")
	beego.Router("/app/table", &controllers.AppController{}, "*:Table")
	beego.Router("/app/add", &controllers.AppController{}, "*:Add")
	beego.Router("/app/ajaxSave", &controllers.AppController{}, "*:AjaxSave")
	beego.Router("/app/ajaxDel", &controllers.AppController{}, "*:AjaxDel")

	beego.Router("/app/update", &controllers.UpdateController{}, "*:List")
	beego.Router("/app/update/add", &controllers.UpdateController{}, "*:Add")
	beego.Router("/app/update/table", &controllers.UpdateController{}, "*:Table")
	beego.Router("/app/update/ajaxSave", &controllers.UpdateController{}, "*:AjaxSave")
	beego.Router("/app/update/edit", &controllers.UpdateController{}, "*:Edit")
	beego.Router("/app/update/ajaxDel", &controllers.UpdateController{}, "*:AjaxDel")

	beego.Router("/app/stop", &controllers.StopController{}, "*:List")
	beego.Router("/app/stop/add", &controllers.StopController{}, "*:Add")
	beego.Router("/app/stop/table", &controllers.StopController{}, "*:Table")
	beego.Router("/app/stop/ajaxSave", &controllers.StopController{}, "*:AjaxSave")
	beego.Router("/app/stop/edit", &controllers.StopController{}, "*:Edit")
	beego.Router("/app/stop/ajaxDel", &controllers.StopController{}, "*:AjaxDel")

	beego.Router("/app/channel", &controllers.ChannelController{}, "*:List")
	beego.Router("/app/channel/add", &controllers.ChannelController{}, "*:Add")
	beego.Router("/app/channel/table", &controllers.ChannelController{}, "*:Table")
	beego.Router("/app/channel/ajaxSave", &controllers.ChannelController{}, "*:AjaxSave")
	beego.Router("/app/channel/edit", &controllers.ChannelController{}, "*:Edit")
	beego.Router("/app/channel/ajaxDel", &controllers.ChannelController{}, "*:AjaxDel")

	beego.Router("/app/advert", &controllers.AdvertController{}, "*:List")
	beego.Router("/app/advert/add", &controllers.AdvertController{}, "*:Add")
	beego.Router("/app/advert/table", &controllers.AdvertController{}, "*:Table")
	beego.Router("/app/advert/ajaxSave", &controllers.AdvertController{}, "*:AjaxSave")
	beego.Router("/app/advert/edit", &controllers.AdvertController{}, "*:Edit")
	beego.Router("/app/advert/ajaxDel", &controllers.AdvertController{}, "*:AjaxDel")

	psw, salt := utils.Password(4, "")

	beego.Info("--------------psw:", psw, " | salt:", salt)
}
