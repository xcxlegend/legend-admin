package models

// 来源model_func模板将替换User为具体model
// User => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func ListUser(limit, offset int) (int64, []*User) {
	query := orm.NewOrm().QueryTable(new(User))
	total, err := query.Count()
	all := []*User{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListUser count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddUser(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateUser(param orm.Params)  (int64, error) {
	return orm.NewOrm().QueryTable(new(User)).Update(param)
}

func DelUser(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelUsers(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(User)).Filter("id__in", ids).Delete()
}
