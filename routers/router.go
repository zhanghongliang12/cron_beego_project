package routers

import (
	"cron_beeproject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//user get
	beego.Router("/user?:id", &controllers.UserController{}, "get")
}
