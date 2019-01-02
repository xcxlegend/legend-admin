package controllers

import (
	"models"
	"encoding/json"
)

type MainController struct {
	AdminController
}

type MenuTree struct{
	models.Node
	Child []*MenuTree
}


func (this *MainController) Index() {
	menu := this.genMenuHtml()
	this.Data["menu"] = menu
	this.Data["user"] = this.getLoginUser()
	this.TplName = "common/frame.html"
}

func (this *MainController) genMenuHtml() string {
	roles := models.AllSortedNodeMenu()
	html, _ := json.MarshalIndent(roles, "","	")
	return string(html)
}

func (this *MainController) Home() {
	this.TplName = "common/home.html"
}
