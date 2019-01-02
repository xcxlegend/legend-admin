package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	. "common"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// RoleUser => CONTROLLER_NAME
// RoleUser => MODEL_NAME


type RoleUserController struct {
	AdminController
}

const (
	RoleUser_DEFAULT_LIMIT = 10
)

type UserInRole struct{
	*models.User
	Selected bool
}

func ListAllRoleUser(rid int) []*UserInRole {
	users := models.AllUsers()
	rusrs := models.AllRoleUsers(rid)
	roleUsers := map[int]struct{}{}
	for _, ru := range rusrs{
		roleUsers[ru.Uid] = struct{}{}
	}

	usrInRole := []*UserInRole{}
	for _, user := range users {
		_, ok := roleUsers[user.Id]
		usrInRole = append(usrInRole, &UserInRole{
			User: user,
			Selected: ok,
		})
	}
	return usrInRole
}


func (this *RoleUserController) Get() {
	id, _ := this.GetInt("id", 0)
	if this.IsAjax() {
		if id == 0 {
			this.Rsp(map[string]interface{}{
				"count": 0,
				"code": 0,
			})
			return
		}
		l := ListAllRoleUser(id)
		this.Rsp(map[string]interface{}{
			"count": len(l),
			"data": l,
			"code": 0,
		})
	} else {
		this.Data["id"] = id
		this.TplName = "role_user/index.html"
	}
}

func (this *RoleUserController) Post() {
	valid := validation.Validation{}
	m := new(models.RoleUser)
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
	if _, err := models.AddRoleUser(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *RoleUserController) Put() {

	rid, _ := this.GetInt("rid")
	if rid == 0 {
		this.RspError("rid error")
		return
	}

	ids := this.GetStrings("uids[]")

	if _, err := models.DeleteRoleUserByRid(rid); err != nil  {
		beego.Error("DeleteRoleUserByRid error:", err, rid)
		this.RspError(err.Error())
		return
	}

	uids := UniqueIntSlice(StringSlice2Int(ids))
	mds := []*models.RoleUser{}
	for _, uid := range uids {
		if uid > 0 {
			mds = append(mds, &models.RoleUser{
				Rid: rid,
				Uid: uid,
			})
		}
	}
	if _, err := models.MultiInsertRoleUser(mds); err != nil {
		beego.Error("MultiInsertRoleUser error:", err, rid)
		this.RspError(err.Error())
		return
	}
	this.RspSuccess("ok")
}

func (this *RoleUserController) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelRoleUsers(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *RoleUserController) Options()  {
	this.Rsp(new(models.RoleUser).Options())
}