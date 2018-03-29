package routers

import (
	"zd112_backstage/controllers"
	"github.com/astaxie/beego"
	"time"
	"github.com/jiangshide/GoComm/utils"
	"strings"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/login", &controllers.PermissionUserController{}, "*:Login")
	beego.Router("/reg", &controllers.PermissionUserController{}, "*:Reg")
	beego.Router("/logOut", &controllers.PermissionUserController{}, "*:LogOut")
	beego.Router("/user/edit", &controllers.PermissionUserController{}, "*:Edit")

	beego.Router("/upload", &controllers.UploadController{}, "*:Upload")
	beego.ErrorController(&controllers.ErrorController{})

	comm()
}

func comm() {
	actionStr := "list,add,edit,table,ajaxSave,ajaxDel"
	router := make(map[string]beego.ControllerInterface, 0)
	router["/permission/user:"+actionStr] = &controllers.PermissionUserController{}
	router["/permission/role:"+actionStr] = &controllers.PermissionRoleController{}
	router["/permission/auth:"+actionStr] = &controllers.PermissionAuthContrller{}
	router["/web/banner:"+actionStr] = &controllers.BannerController{}
	router["/web/nav:"+actionStr] = &controllers.NavController{}
	router["/web/nav/university:"+actionStr] = &controllers.UniversityController{}
	router["/nation:"+actionStr] = &controllers.NationController{}
	router["/areaCode:"+actionStr] = &controllers.AreaCodeController{}
	router["/area:"+actionStr] = &controllers.AreaController{}

	beego.Router("/nation/list", &controllers.NationController{}, "*:List")
	for k, v := range router {
		kArr := strings.Split(k, ":")
		actions := strings.Split(kArr[1], ",")
		for _, action := range actions {
			rootPath := kArr[0]
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
