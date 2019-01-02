package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// Node => CONTROLLER_NAME
// Node => MODEL_NAME


type NodeController struct {
	AdminController
}

const (
	Node_DEFAULT_LIMIT = 10
)


func (this *NodeController) Get() {
	if this.IsAjax() {
		limit, _ := this.GetInt("limit", Node_DEFAULT_LIMIT)
		page, _ := this.GetInt("page", 1)
		if page <= 0 {
			page = 1
		}
		offset := (page-1) * limit
		total, list := models.ListNode(limit, offset)
		this.Rsp(map[string]interface{}{
			"count": total,
			"data": list,
			"code": 0,
		})
	} else {
		this.TplName = "node/index.html"
	}
}

func (this *NodeController) Post() {
	valid := validation.Validation{}
	m := new(models.Node)
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
	if _, err := models.AddNode(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *NodeController) Put() {

	valid := validation.Validation{}
	m := new(models.Node)
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
	if _, err := models.UpdateNode(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *NodeController) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelNodes(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *NodeController) Options()  {
	this.Rsp(new(models.Node).Options())
}