package models

// 来源model_func模板将替换Node为具体model
// Node => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


func ListNode(limit, offset int) (int64, []*Node) {
	query := orm.NewOrm().QueryTable(new(Node))
	total, err := query.Count()
	all := []*Node{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListNode count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddNode(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateNode(m interface{})  (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func DelNode(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelNodes(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(Node)).Filter("id__in", ids).Delete()
}
