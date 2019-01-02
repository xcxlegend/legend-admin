package controllers

import (
	//"github.com/astaxie/beego"
	"models"
	"github.com/pkg/errors"
	. "common"
)


type PublicController struct {
	BaseController
}

func (this *PublicController) Login() {
	if this.IsAjax(){
		if err := this.checkLoginAndAuth(); err != nil {
			this.RspError(err.Error())
		} else {
			this.RspSuccessWithMsg("", "ok")
		}
	} else {
		this.TplName = "public/login.html"
	}
}

func (this *PublicController) Logout()  {
	this.DelSession("user")
	this.DelSession("acc")
	this.Redirect("/public/login", 301)
}

// 判断登陆和授权
func (this *PublicController) checkLoginAndAuth() error {
	username := this.GetString("username")
	password := this.GetString("password")
	user := models.GetUserByUsername(username)
	if user == nil {
		return errors.New("用户不存在")
	}

	if user.State != 1 {
		return errors.New("用户被禁用")
	}

	if user.Password != models.HashPassword(password) {
		return errors.New("密码错误")
	}

	this.SetSession("user", user)
	acc := this.GetRoleNodeList(user.Id)
	this.SetSession("acc", acc)
	return nil
}

func (this *PublicController) GetRoleNodeList(uid int) AccLists {
	role_ids := models.GetRidsByUid(uid)
	nodes := models.GetNodesByRids(role_ids)
	acc := AccLists{}
	for _, node := range nodes {
		if _, ok := acc[node.Router]; !ok {
			acc[node.Router] = map[string]struct{}{}
		}
		acc[node.Router][node.RouterType] = struct{}{}
	}
	return acc
}