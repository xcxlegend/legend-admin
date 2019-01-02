package models

// 来源model_func模板将替换Role为具体model
// Role => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func ListRole(limit, offset int) (int64, []*Role) {
	query := orm.NewOrm().QueryTable(new(Role))
	total, err := query.Count()
	all := []*Role{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListRole count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddRole(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateRole(m interface{})  (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func DelRole(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelRoles(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(Role)).Filter("id__in", ids).Delete()
}
