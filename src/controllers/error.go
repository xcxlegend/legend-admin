package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error401()  {
	this.TplName = "error/401.html"
}