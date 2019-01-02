package models

// 来源model_func模板将替换T__MODELTEMPLDATE__为具体model
// T__MODELTEMPLDATE__ => MODEL_NAME

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type T__MODELTEMPLDATE__ struct {Base}
func ListT__MODELTEMPLDATE__(limit, offset int) (int64, []*T__MODELTEMPLDATE__) {
	query := orm.NewOrm().QueryTable(new(T__MODELTEMPLDATE__))
	total, err := query.Count()
	all := []*T__MODELTEMPLDATE__{}
	if err != nil {
		beego.Error(fmt.Sprintf(`ListT__MODELTEMPLDATE__ count err:%v`, err))
		return total, all
	}
	query.Limit(limit).Offset(offset).All(&all)
	return total, all
}

func AddT__MODELTEMPLDATE__(m interface{}) (int64, error){
	o := orm.NewOrm()
	return o.Insert(m)
}

func UpdateT__MODELTEMPLDATE__(m interface{})  (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func DelT__MODELTEMPLDATE__(m interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(m)
}

func DelT__MODELTEMPLDATE__s(ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	return orm.NewOrm().QueryTable(new(T__MODELTEMPLDATE__)).Filter("id__in", ids).Delete()
}
