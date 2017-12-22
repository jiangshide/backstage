package routers

import (
	"zd112_backstage/controllers"
	"github.com/astaxie/beego"
	"time"
	"strings"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/reg", &controllers.UserController{}, "*:Reg")
	beego.Router("/admin", &controllers.UserController{}, "*:Admin")
	beego.Router("/loginOut", &controllers.UserController{}, "*:LoginOut")
	beego.Router("/backstage", &controllers.BackstageController{}, "*:Index")
	beego.Router("/backstage/home/start", &controllers.BackstageController{}, "*:Start")
	beego.Router("/backstage/user/edit", &controllers.UserController{}, "*:Edit")

	beego.Router("/backstage/admin/list", &controllers.AdminController{}, "*:List")
	beego.Router("/backstage/admin/add", &controllers.AdminController{}, "*:Add")
	beego.Router("/backstage/admin/edit", &controllers.AdminController{}, "*:Edit")
	beego.Router("/backstage/admin/table", &controllers.AdminController{}, "*:Table")

	beego.AutoRouter(&controllers.AdminController{})

	beego.Router("/backstage/role/list", &controllers.RoleController{}, "*:List")
	beego.Router("/backstage/role/add", &controllers.RoleController{}, "*:Add")
	beego.Router("/backstage/role/edit", &controllers.RoleController{}, "*:Edit")
	beego.Router("/backstage/role/table", &controllers.RoleController{}, "*:Table")
	beego.AutoRouter(&controllers.RoleController{})

	beego.Router("/backstage/auth/list", &controllers.AuthController{}, "*:List")
	//beego.Router("/backstage/auth/add",&controllers.AuthController{},"*:Add")
	//beego.Router("/backstage/auth/edit",&controllers.AuthController{},"*:Edit")
	//beego.Router("/backstage/auth/table",&controllers.AuthController{},"*:Table")
	beego.AutoRouter(&controllers.AuthController{})

	commRouter()

	beego.Router("/upload", &controllers.BaseController{}, "*:Upload")
	beego.ErrorController(&controllers.ErrorController{})

	/**
	 the api
	 */
	beego.Router("/api/user", &controllers.UserApiController{}, "*:Login")
	beego.Router("/api/user", &controllers.UserApiController{}, "*:Reg")
	beego.Router("/api/user", &controllers.UserApiController{}, "*:List")
	beego.Router("/api/splash", &controllers.SpashApiController{}, "*:Splash")
}

func commRouter() {
	actionStr := "list,add,edit,table,ajaxSave,ajaxDel"
	router := make(map[string]beego.ControllerInterface, 0)
	router["nation:"+actionStr] = &controllers.NationController{}
	router["area/continent:"+actionStr] = &controllers.ContinentController{}
	router["area/state:"+actionStr] = &controllers.StateController{}
	router["area/province:"+actionStr] = &controllers.ProvinceController{}
	router["area/city:"+actionStr] = &controllers.CityController{}
	router["area/region:"+actionStr] = &controllers.RegionController{}
	router["area/county:"+actionStr] = &controllers.CountyController{}
	router["area/town:"+actionStr] = &controllers.TownController{}
	router["area/country:"+actionStr] = &controllers.CountryController{}
	router["area/village:"+actionStr] = &controllers.VillageController{}
	router["area/group:"+actionStr] = &controllers.GroupController{}
	router["area/team:"+actionStr] = &controllers.TeamController{}
	router["web/banner:"+actionStr] = &controllers.BannerController{}
	router["tools/qrcode:"+actionStr] = &controllers.QrcodeController{}
	router["tools/compress:"+actionStr] = &controllers.CompressController{}
	router["tools/formattype:"+actionStr] = &controllers.FormatTypeController{}
	router["tools/format:"+actionStr] = &controllers.FormatController{}
	router["app/channel:"+actionStr] = &controllers.ChannelController{}
	router["app/app_name:"+actionStr] = &controllers.AppNameController{}
	router["app/pkgs:"+actionStr] = &controllers.PkgsController{}
	router["app/version:"+actionStr] = &controllers.VersionController{}
	router["app/code:"+actionStr] = &controllers.CodeController{}
	router["app/env:"+actionStr] = &controllers.EnvController{}
	router["app/build:"+actionStr] = &controllers.BuildController{}
	router["app/type:"+actionStr] = &controllers.TypeController{}
	router["test/app:"+actionStr] = &controllers.AppController{}
	router["test/environment:"+actionStr] = &controllers.EnvironmnetController{}
	router["test/project:"+actionStr] = &controllers.ProjectController{}
	router["test/test:"+actionStr] = &controllers.TestController{}
	router["api/key:"+actionStr] = &controllers.KeyController{}
	for k, v := range router {
		kArr := strings.Split(k, ":")
		path := "/backstage/" + kArr[0]
		actions := strings.Split(kArr[1], ",")
		for _, action := range actions {
			rootPath := path
			if action != "list" {
				rootPath += "/" + strings.ToLower(action)
			}
			action = "*:" + utils.StrFirstToUpper(action)
			//beego.Info("rootPath:", rootPath, " | action:", action)
			beego.Router(rootPath, v, action)
		}
	}
}

func taskTime() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			//res, err := utils.ExecCommand("/usr/bin/git --git-dir=" + utils.GetCurrentDir("") + "/.git checkout master")
			//beego.Info("res:", string(res), " | err:", err)
		}
	}()
}
