package models

import "github.com/astaxie/beego/orm"

type Role struct{
	Base
	Id     	int		`orm:"size(11)"`
	Title 	string	`orm:"size(32)" form:"Title,text,角色名" valid:"Required;MaxSize(32);MinSize(1)"`
	Remark 	string	`orm:"size(32)" form:"Remark,textarea,备注"`
	State 	int8	`form:"State,select,状态,Options:StateSelector"`
}

// 需要一个init方法注册相应的结构体 会自动变成小写的表名 如user
func init()  {
	// 需要注册model, 在gen的时候使用name来索引相应的结构体
	registerBuildModel("role", new(Role))
	orm.RegisterModel(new(Role))
}