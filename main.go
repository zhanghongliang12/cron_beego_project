package main

import (
	_ "cron_beeproject/models"
	_ "cron_beeproject/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 根据环境判断是否走swagger 接口
	runmode, _ := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true

		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run("127.0.0.1:8080")

}
