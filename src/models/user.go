package models

import (
	"github.com/astaxie/beego/orm"
	"common"
)

// 模板例子  需要如此的一个struct 定义orm, form, valid的tag 分别对应mysql的结构 表单和验证相关
// form:"Name,type,Label()" type: text,password,select,date,time,datetime; 日期会自动转为timestamp请使用int size(11), 不设置size
type User struct{
	Base
	Id            int
	Username      string    `orm:"unique;size(32)" form:"Username,text,用户名" valid:"Required;MaxSize(32);MinSize(3)"`
	Password      string    `orm:"size(32)" form:"Password,password,密码" json:"-" valid:"MaxSize(32)"`
	Nickname      string    `orm:"unique;size(32)" form:"Nickname,text,昵称" valid:"Required;MaxSize(32);MinSize(3)"`
	State 		  int8		`form:"State,select,状态,Options:StateSelector"`
}

// 需要一个init方法注册相应的结构体 会自动变成小写的表名 如user
func init()  {
	// 需要注册model, 在gen的时候使用name来索引相应的结构体
	registerBuildModel("user", new(User))
	orm.RegisterModel(new(User))
}

func (this *User) Options() []map[string]string {
	return []map[string]string{
		{"key": "1", "value": "1"},
	}
}

func (this *User) PreAdd()  {
	this.Password = HashPassword(this.Password)
}

func (this *User) PreUpdate() {
	if this.Password != ""{
		this.Password = HashPassword(this.Password)
	}
}

func HashPassword(password string) string  {
	return common.Md5Str(password + "LEGENDADMIN")
}

func GetUserByUsername(username string) *User {
	u := &User{Username: username}
	if orm.NewOrm().Read(u, "username") != nil {
		return nil
	}
	return u
}

func AllUsers() []*User {
	users := []*User{}
	orm.NewOrm().QueryTable(new(User)).All(&users)
	return users
}