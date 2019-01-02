package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/public/login", &controllers.PublicController{}, "*:Login")
    beego.Router("/public/logout", &controllers.PublicController{}, "*:Logout")
	beego.Router("/admin/index", &controllers.MainController{}, "*:Index")
	beego.Router("/admin/home", &controllers.MainController{}, "*:Home")
}
