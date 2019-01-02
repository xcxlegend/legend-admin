package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// User => CONTROLLER_NAME
// User => MODEL_NAME


type UserController struct {
	AdminController
}

const (
	User_DEFAULT_LIMIT = 10
)


func (this *UserController) Get() {
	if this.IsAjax() {
		limit, _ := this.GetInt("limit", User_DEFAULT_LIMIT)
		page, _ := this.GetInt("page", 1)
		if page <= 0 {
			page = 1
		}
		offset := (page-1) * limit
		total, list := models.ListUser(limit, offset)
		this.Rsp(map[string]interface{}{
			"count": total,
			"data": list,
			"code": 0,
		})
	} else {
		this.TplName = "user/index.html"
	}
}

func (this *UserController) Post() {
	valid := validation.Validation{}
	m := new(models.User)
	this.ParseForm(m)
	b, err := valid.Valid(m)
	if err != nil {
		this.RspError( err.Error())
		return
	}
	if !b {
		for _, e := range valid.Errors{
			this.RspError( e.Key + ": " + e.Message)
			return
		}
	}
	m.PreAdd()
	if _, err := models.AddUser(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *UserController) Put() {

	valid := validation.Validation{}
	m := new(models.User)
	this.ParseForm(m)
	b, err := valid.Valid(m)
	if err != nil {
		this.RspError( err.Error())
		return
	}
	if !b {
		for _, e := range valid.Errors{
			this.RspError( e.Key + ": " + e.Message)
			return
		}
	}
	m.PreUpdate()

	param := orm.Params{}
	if m.Username != ""{
		param["Username"] = m.Username
	}

	if m.Nickname != ""{
		param["Nickname"] = m.Nickname
	}

	if m.Password != ""{
		param["Password"] = m.Password
	}

	param["State"] = m.State

	if _, err := models.UpdateUser(param); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *UserController) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelUsers(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *UserController) Options()  {
	this.Rsp(new(models.User).Options())
}