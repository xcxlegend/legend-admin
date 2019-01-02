package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RoleNode struct{
	Base
	Id     	int		`orm:"size(11)"`
	Rid 	int		`orm:"size(11)"` // role id
	Nid 	int		`orm:"size(11)"` // node_id
}

var RoleNodeQueryTableInst *RoleNode


// 需要一个init方法注册相应的结构体 会自动变成小写的表名 如user
func init()  {
	RoleNodeQueryTableInst = new(RoleNode)
	// 需要注册model, 在gen的时候使用name来索引相应的结构体
	registerBuildModel("role_node", RoleNodeQueryTableInst)
	orm.RegisterModel(RoleNodeQueryTableInst)
}

func AllRoleNodes(rid int) []*RoleNode {
	rusrs := []*RoleNode{}
	orm.NewOrm().QueryTable(RoleNodeQueryTableInst).Filter("rid", rid).All(&rusrs)
	return rusrs
}

func DeleteRoleNodeByRid(rid int) (int64, error) {
	return orm.NewOrm().QueryTable(RoleNodeQueryTableInst).Filter("rid", rid).Delete()
}

func MultiInsertRoleNode(mds []*RoleNode) (int64, error) {
	return orm.NewOrm().InsertMulti(0, mds)
}

// 根据rids列表获取全部的node
func GetNodesByRids(rids []int) []*Node {
	prls := orm.ParamsList{}
	if _, err := orm.NewOrm().QueryTable(RoleNodeQueryTableInst).Filter("rid__in", rids).Distinct().ValuesFlat(&prls, "nid"); err != nil {
		beego.Error("rid in rids:", rids, err)
		return []*Node{}
	}
	nids := []int{}

	for _, p := range prls{
		nid, _ := p.(int64)
		nids = append(nids, int(nid))
	}

	return GetNodesByIds(nids)
}