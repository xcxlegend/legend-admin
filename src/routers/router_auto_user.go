package routers


import (
	 "controllers"
	 "github.com/astaxie/beego"
)

func init() {
	 beego.Router("/admin/user", &controllers.UserController{})
}
