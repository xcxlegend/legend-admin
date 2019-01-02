package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// Role => CONTROLLER_NAME
// Role => MODEL_NAME


type RoleController struct {
	AdminController
}

const (
	Role_DEFAULT_LIMIT = 10
)


func (this *RoleController) Get() {
	if this.IsAjax() {
		limit, _ := this.GetInt("limit", Role_DEFAULT_LIMIT)
		page, _ := this.GetInt("page", 1)
		if page <= 0 {
			page = 1
		}
		offset := (page-1) * limit
		total, list := models.ListRole(limit, offset)
		this.Rsp(map[string]interface{}{
			"count": total,
			"data": list,
			"code": 0,
		})
	} else {
		this.TplName = "role/index.html"
	}
}

func (this *RoleController) Post() {
	valid := validation.Validation{}
	m := new(models.Role)
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
	if _, err := models.AddRole(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *RoleController) Put() {

	valid := validation.Validation{}
	m := new(models.Role)
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
	if _, err := models.UpdateRole(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *RoleController) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelRoles(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *RoleController) Options()  {
	this.Rsp(new(models.Role).Options())
}

