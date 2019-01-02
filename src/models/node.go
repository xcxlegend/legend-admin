package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

/**
	mysql table
	id router title rid sort state
 */

type Node struct{
	Base
	Id     	int		`orm:"size(11)"`
	Router 	string	`orm:"size(50)" form:"Router,text,路由"`
	RouterType string `orm:"size(10)" form:"RouterType,select,访问类型,Options:RouterTypeSelector"`
	Title 	string	`orm:"size(32)" form:"Title,text,标题" valid:"Required;MaxSize(32);MinSize(1)"`
	Rid 	int		`form:"Rid,select,上级ID,OptionsUrl:/admin/node"` // role pid
	Sort 	int8	`form:"Sort,text,排序"`
	State 	int8	`form:"State,select,状态,Options:StateSelector"`
	Icon	string	`orm:"size(32)" form:"Icon,text,图标"`
	IsMenu  int8	`form:"IsMenu,select,是否菜单,Options:YesOrNoSelector"`
}

func init()  {
	// 需要注册model, 在gen的时候使用name来索引相应的结构体
	registerBuildModel("node", new(Node))
	orm.RegisterModel(new(Node))
}

func (this *Node) Options() []map[string]string {
	roles := []*Node{}
	orm.NewOrm().QueryTable(this).All(&roles, "id", "title")
	options := []map[string]string{}
	for _, r := range roles{
		options = append(options, map[string]string{
			"key": fmt.Sprintf(`%v`, r.Id), "value": r.Title,
		})
	}
	return options
}

func AllSortedNodeMenu() []*Node {
	roles := []*Node{}
	orm.NewOrm().QueryTable(new(Node)).Filter("state", 1).Filter("is_menu", 1).OrderBy("sort").All(&roles)
	return roles
}


func AllSortedNode() []*Node {
	roles := []*Node{}
	orm.NewOrm().QueryTable(new(Node)).Filter("state", 1).OrderBy("sort").All(&roles)
	return roles
}

func GetNodesByIds(ids []int) []*Node{
	nodes := []*Node{}
	if _, err := orm.NewOrm().QueryTable(new(Node)).Filter("id__in", ids).Filter("state", 1).All(&nodes); err != nil {
		beego.Error("GetNodesByIds:", ids, err)
	}
	return nodes
}