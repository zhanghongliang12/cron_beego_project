package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}
type name_test struct {
	name string
	age  int
}

func (self *UserController) Get() {
	id, err := self.GetInt(":id")
	if err != nil {
		fmt.Println("内部参与转换类型失败")
		return
	}
	fmt.Println("id:---------->", id)

	var name name_test
	name.name = "张三"
	name.age = 18
	fmt.Println(name)
}
