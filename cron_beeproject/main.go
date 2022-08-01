package main

import (
	_ "cron_beeproject/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "cron_beeproject/models"

)

func main() {
	beego.Run()
}

