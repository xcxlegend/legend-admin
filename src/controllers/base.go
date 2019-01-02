package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare()  {
	this.Layout = "common/layout.html"
}

func (this *BaseController) Rsp(r interface{}) {
	this.Data["json"] = r
	this.ServeJSON()
}

func (this *BaseController) RspSuccess(data interface{}) {
	this.Rsp(map[string]interface{}{
		"data": data,
		"error": "",
	})
}

func (this *BaseController) RspSuccessWithMsg(data interface{}, msg string) {
	this.Rsp(map[string]interface{}{
		"data": data,
		"error": "",
		"msg": msg,
	})
}

func (this *BaseController) RspError(err string) {
	this.Rsp(map[string]interface{}{
		"error": err,
	})
}

func (this *BaseController) ParseJSONIdArray() (data []int64) {
	data = []int64{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	return
}