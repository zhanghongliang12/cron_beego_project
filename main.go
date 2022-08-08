package main

import (
	_ "cron_beeproject/models"
	_ "cron_beeproject/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run("127.0.0.1:8080")
}
