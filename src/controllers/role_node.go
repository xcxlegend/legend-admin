package controllers

import (
	"models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	. "common"
)

// 来源controller模板将替换T__CONTROLLERNAME为具体controller
// RoleNode => CONTROLLER_NAME
// RoleNode => MODEL_NAME


type RoleNodeController struct {
	AdminController
}

const (
	RoleNode_DEFAULT_LIMIT = 10
)


type NodeInRole struct{
	*models.Node
	Selected bool
}

func ListAllRoleNode(rid int) []*NodeInRole {
	nodes := models.AllSortedNode()
	rnode := models.AllRoleNodes(rid)
	roleNodes := map[int]struct{}{}
	for _, rn := range rnode{
		roleNodes[rn.Nid] = struct{}{}
	}

	nodeInRole := []*NodeInRole{}
	for _, node := range nodes {
		_, ok := roleNodes[node.Id]
		nodeInRole = append(nodeInRole, &NodeInRole{
			Node: node,
			Selected: ok,
		})
	}
	return nodeInRole
}


func (this *RoleNodeController) Get() {
	id, _ := this.GetInt("id", 0)
	if this.IsAjax() {
		if id == 0 {
			this.Rsp(map[string]interface{}{
				"count": 0,
				"code": 0,
			})
			return
		}
		l := ListAllRoleNode(id)
		this.Rsp(map[string]interface{}{
			"count": len(l),
			"data": l,
			"code": 0,
		})
	} else {
		this.Data["id"] = id
		this.TplName = "role_node/index.html"
	}
}

func (this *RoleNodeController) Post() {
	valid := validation.Validation{}
	m := new(models.RoleNode)
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
	if _, err := models.AddRoleNode(m); err != nil {
		this.RspError( err.Error())
		return
	}
	this.RspSuccess(m)
	return
}

func (this *RoleNodeController) Put() {

	rid, _ := this.GetInt("rid")
	if rid == 0 {
		this.RspError("rid error")
		return
	}

	ids := this.GetStrings("nids[]")

	if _, err := models.DeleteRoleNodeByRid(rid); err != nil  {
		beego.Error("DeleteRoleNodeByRid error:", err, rid)
		this.RspError(err.Error())
		return
	}

	nids := UniqueIntSlice(StringSlice2Int(ids))
	mds := []*models.RoleNode{}
	for _, uid := range nids{
		if uid > 0 {
			mds = append(mds, &models.RoleNode{
				Rid: rid,
				Nid: uid,
			})
		}
	}
	if _, err := models.MultiInsertRoleNode(mds); err != nil {
		beego.Error("MultiInsertRoleNode error:", err, rid)
		this.RspError(err.Error())
		return
	}
	this.RspSuccess("ok")
}

func (this *RoleNodeController) Delete() {
	ids := this.ParseJSONIdArray()
	if _, err := models.DelRoleNodes(ids); err != nil {
		this.RspError(err.Error())
		return
	}
	this.RspSuccessWithMsg(map[string]interface{}{
		"id": ids,
	}, "ok")
	return
}

func (this *RoleNodeController) Options()  {
	this.Rsp(new(models.RoleNode).Options())
}