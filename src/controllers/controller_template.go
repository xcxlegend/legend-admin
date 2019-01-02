package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// T__CONTROLLERNAME__ => CONTROLLER_NAME
// T__MODELTEMPLDATE__ => MODEL_NAME


type T__CONTROLLERNAME__Controller struct {
	AdminController
}

const (
	T__CONTROLLERNAME___DEFAULT_LIMIT = 10
)


func (this *T__CONTROLLERNAME__Controller) Get() {
	if this.IsAjax() {
		limit, _ := this.GetInt("limit", T__CONTROLLERNAME___DEFAULT_LIMIT)
		page, _ := this.GetInt("page", 1)
		if page <= 0 {
			page = 1
		}
		offset := (page-1) * limit
		total, list := models.ListT__MODELTEMPLDATE__(limit, offset)
		this.Rsp(map[string]interface{}{
			"count": total,
			"data": list,
			"code": 0,
		})
	} else {
		this.TplName = "__VIEW_INDEX__"
	}
}

func (this *T__CONTROLLERNAME__Controller) Post() {
	valid := validation.Validation{}
	m := new(models.T__MODELTEMPLDATE__)
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
	if _, err := models.AddT__MODELTEMPLDATE__(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *T__CONTROLLERNAME__Controller) Put() {

	valid := validation.Validation{}
	m := new(models.T__MODELTEMPLDATE__)
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
	if _, err := models.UpdateT__MODELTEMPLDATE__(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *T__CONTROLLERNAME__Controller) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelT__MODELTEMPLDATE__s(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *T__CONTROLLERNAME__Controller) Options()  {
	this.Rsp(new(models.T__MODELTEMPLDATE__).Options())
}