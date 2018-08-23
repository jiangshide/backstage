package models

import (
	"github.com/astaxie/beego/orm"
)

type Update struct {
	Id         int64
	Name       string //应用名称:
	Platform   int    //平台:0~无限制, 1~android, 2~ios, 3~web
	Channel    string //渠道
	Version    string //版本
	Content    string //更新内容提示
	Url        string //下载地址
	Status     int    //更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新
	Count      int64  //更新次数
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Update) TableName() string {
	return TableName("app_update")
}

func (this *Update) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Update) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Update) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Update) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName()).OrderBy("-id")
}

func (this *Update) List(pageSize, offSet int) (list []*Update, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Stop struct {
	Id         int64
	Name       string //名称:
	Channel    string //渠道名称:
	Version    string //版本:3.0.0.18
	Url        string //app本地地址
	Platform   int    //平台:0~无限制, 1~android, 2~ios, 3~web
	Count      int64
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Stop) TableName() string {
	return TableName("app_stop")
}

func (this *Stop) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Stop) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Stop) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Stop) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName()).OrderBy("-id")
}

func (this *Stop) List(pageSize, offSet int) (list []*Stop, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type App struct {
	Id          int64
	Name        string //应用名称:
	Pkg         string //包名:
	FriendId    string //Friend_ID:
	Channel     string //渠道名称:
	Version     string //版本:3.0.0.18
	VersionCode int    //版本号:3019
	Environment string //环境:production,preproduction,stage1,stage2,stage3
	Status      int    //状态:1:打包中,2:打包成功,3:打包失败,4,无
	Cmd         string //打包脚本
	Log         string //打包日志详情
	ErrMsg      string //错误信息
	QrUrl       string //本地二维码:可扫描下载
	AppUrl      string //app本地地址
	Platform    int    //平台:0~无限制, 1~android, 2~ios, 3~web
	Size        int64  //文件尺寸
	CreateId    int64
	UpdateId    int64
	CreateTime  int64
	UpdateTime  int64
}

func (this *App) TableName() string {
	return TableName("app")
}

func (this *App) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *App) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *App) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *App) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *App) List(pageSize, offSet int) (list []*App, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Channel struct {
	Id         int64
	Name       string //应用名称:
	Pkg        string //包名:
	FriendId   string //Friend_ID:
	Channel    string //渠道名称:
	Platform   int    //平台:0~无限制, 1~android, 2~ios, 3~web
	Status     int    //状态:1:正常,2:下架,3:异常
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Channel) TableName() string {
	return TableName("app_channel")
}

func (this *Channel) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Channel) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Channel) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Channel) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName()).OrderBy("-id")
}

func (this *Channel) List(pageSize, offSet int) (list []*Channel, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Advert struct {
	Id         int64
	Name       string //广告名称
	Type       int    //广告类型
	Platform   int    //平台:0~无限制, 1~android, 2~ios, 3~web
	Price      int    //价格
	Times      int    //展示次数
	Path       string //内容绝对路径
	Url        string //内容地址
	Status     int    //状态:1:正常,2:下架,3:异常
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Advert) TableName() string {
	return TableName("app_advert")
}

func (this *Advert) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Advert) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Advert) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Advert) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Advert) List(pageSize, offSet int) (list []*Advert, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
