package controllers

import (
	"fmt"
	"time"
	"github.com/pkg/errors"
	"github.com/astaxie/beego"
	"strings"
	"github.com/jiangshide/GoComm/utils"
)

type BaseApiController struct {
	BaseController
	code    int
	msg     interface{}
	content interface{}
	isToken bool
}

const (
	RES_OK           = 0
	FALSE            = -1
	TOKEN_INVALIDATE = -2
)

func (this *BaseApiController) Prepare() {
	this.isApi = true
	this.version = beego.AppConfig.String("version.api")
	this.prepare()
	if _, err := this.checkToken(); err != nil {
		this.code = TOKEN_INVALIDATE
		this.msg = err
		this.response()
	}
}

func (this *BaseApiController) header(key string) string {
	return this.Ctx.Request.Header.Get(key)
}

func (this *BaseApiController) checkToken() (map[string]interface{}, error) {
	auth := this.header("auth")
	beego.Info("auth:", auth)
	if len(auth) == 0 {
		return nil, errors.New("auth is null")
	}
	return utils.UnToken(auth, this.loginName)
}

func (this *BaseApiController) response() {
	res := make(map[string]interface{}, 0)
	res["code"] = this.code
	res["version"] = this.version
	res["date"] = time.Now().Unix()
	if this.code == RES_OK {
		res["res"] = this.content
		if this.isToken {
			tokenMap := make(map[string]interface{}, 0)
			tokenMap["code"] = this.code
			tokenMap["version"] = this.version
			tokenMap["user"] = this.user
			if token, err := utils.Token(tokenMap, this.loginName, time.Now().Add(time.Hour * 1).Unix()); err == nil {
				res["token"] = token
			}
		}
	} else {
		res["msg"] = fmt.Sprint(this.msg)
	}
	this.Data["json"] = res
	this.ServeJSON()
	//this.StopRun()
}

func (this *BaseApiController) getString(key, tips string, minSize int) string {
	value := strings.TrimSpace(this.GetString(key, ""))
	errorMsg := ""
	if len(value) == 0 {
		errorMsg = tips
	} else if len(value) < minSize {
		errorMsg = "长度不能小于" + fmt.Sprint(minSize) + "字符"
	}
	if errorMsg != "" {
		this.code = -1
		this.msg = errorMsg
		this.response()
	}
	return value
}
