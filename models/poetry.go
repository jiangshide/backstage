package models

type Poetry struct {
	Id     int64
	Name   string //诗词名称
	Type   string //类型
	Author string //作者
	Age    string //时代
	Form   string //形式
	From   string //来自于

	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

type Comment struct {
	Id         int64
	Name       string //标题
	Content    string //内容
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}
