package routers


import (
	 "controllers"
	 "github.com/astaxie/beego"
)

func init() {
	 beego.Router("/admin/role_node", &controllers.RoleNodeController{})
}
