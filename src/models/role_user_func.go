package models

// 来源model_func模板将替换RoleUser为具体model
// RoleUser => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func ListRoleUser(limit, offset int) (int64, []*RoleUser) {
	query := orm.NewOrm().QueryTable(new(RoleUser))
	total, err := query.Count()
	all := []*RoleUser{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListRoleUser count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddRoleUser(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateRoleUser(m interface{})  (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func DelRoleUser(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelRoleUsers(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(RoleUser)).Filter("id__in", ids).Delete()
}
