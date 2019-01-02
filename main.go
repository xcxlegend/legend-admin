package main

import (
	_ "routers"
	"github.com/astaxie/beego"
	. "widgets"
	_ "common"
	"controllers"
)
 
func main() {
	beego.BConfig.CopyRequestBody = true
	if !GetWidgetMgr().Run() {
		beego.Warn("widgets run over&stop")
		return
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

