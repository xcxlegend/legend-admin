package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RoleUser struct{
	Base
	Id     	int		`orm:"size(11)"`
	Rid 	int		`orm:"size(11)"` // role id
	Uid 	int		`orm:"size(11)"` // uid
}

// 这个用作于orm QueryTable 的单例
var RoleUserQueryTableInst *RoleUser

// 需要一个init方法注册相应的结构体 会自动变成小写的表名 如user
func init()  {
	// 需要注册model, 在gen的时候使用name来索引相应的结构体
	RoleUserQueryTableInst = new(RoleUser)
	registerBuildModel("role_user", RoleUserQueryTableInst)
	orm.RegisterModel(RoleUserQueryTableInst)
}

func AllRoleUsers(rid int) []*RoleUser {
	rusrs := []*RoleUser{}
	orm.NewOrm().QueryTable(RoleUserQueryTableInst).Filter("rid", rid).All(&rusrs)
	return rusrs
}

func DeleteRoleUserByRid(rid int) (int64, error) {
	return orm.NewOrm().QueryTable(RoleUserQueryTableInst).Filter("rid", rid).Delete()
}

func MultiInsertRoleUser(mds []*RoleUser) (int64, error) {
	return orm.NewOrm().InsertMulti(0, mds)
}

// 根据用户id反查所有的角色id
func GetRidsByUid(uid int) []int {
	prls := orm.ParamsList{}
	ids := []int{}
	if _, err := orm.NewOrm().QueryTable(RoleUserQueryTableInst).Filter("uid", uid).ValuesFlat(&prls, "rid"); err != nil{
		beego.Error("GetRidsByUid:", err)
		return ids
	}

	rprls := orm.ParamsList{}
	if _, err := orm.NewOrm().QueryTable(new(Role)).Filter("id__in", prls).ValuesFlat(&rprls, "id"); err != nil {
		beego.Error("Role id in:", prls, err)
		return ids
	}

	for _, i := range prls{
		id, _ := i.(int64)
		ids = append(ids, int(id))
	}
	return ids
}