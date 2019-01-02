package models

// 来源model_func模板将替换RoleNode为具体model
// RoleNode => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func ListRoleNode(limit, offset int) (int64, []*RoleNode) {
	query := orm.NewOrm().QueryTable(new(RoleNode))
	total, err := query.Count()
	all := []*RoleNode{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListRoleNode count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddRoleNode(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateRoleNode(m interface{})  (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func DelRoleNode(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelRoleNodes(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(RoleNode)).Filter("id__in", ids).Delete()
}
