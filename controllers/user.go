package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

)

type UserController struct {
	beego.Controller
}

func (self *UserController) Get() {
	self.GetInt("")

}