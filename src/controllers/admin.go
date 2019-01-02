package controllers

import (
	. "common"
	"strings"
	"models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Prepare()  {
	this.BaseController.Prepare()
	if beego.BConfig.RunMode != "dev" {
		this.checkAuth()
	}
}

func (this *AdminController) checkAuth()  {
	user := this.GetSession("user")
	if user == nil {
		this.Ctx.Redirect(302, "/public/login")
		return
	}

	// 检查是否是admin权限
	User, _ := user.(*models.User)
	if User.Username == beego.AppConfig.String("accessadmin") {
		return
	}

	acc := this.GetSession("acc")
	if acc == nil {
		this.showNoAccess()
		return
	}

	accList, _ := acc.(AccLists)
	router := strings.Split(this.Ctx.Request.RequestURI, "?")[0]
	r, ok := accList[router]
	if !ok {
		this.showNoAccess()
		return
	}
	if _, ok := r["ALL"]; ok {
		return
	}
	if _, ok := r[this.Ctx.Input.Method()]; ok {
		return
	}
	this.showNoAccess()
}

func (this *AdminController) showNoAccess()  {
	if this.IsAjax(){
		this.RspError("no access")
		return
	}
	this.Abort("401")
}

func (this *AdminController) getLoginUser() models.User  {
	u := this.GetSession("user")
	user, _ := u.(*models.User)
	if user == nil {
		return models.User{}
	}else {
		return *user
	}
}