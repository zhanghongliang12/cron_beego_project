package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	name, _ := beego.AppConfig.String("name")
	db_name, _ := beego.AppConfig.String("db_name")
	user, _ := beego.AppConfig.String("user")
	password, _ := beego.AppConfig.String("password")
	port, _ := beego.AppConfig.String("port")
	host, _ := beego.AppConfig.String("host")
	fmt.Println("-------------")
	fmt.Println(name)
	fmt.Println(db_name)
	fmt.Println(user)
	fmt.Println(password)
	fmt.Println(port)
	fmt.Println(host)
	fmt.Println("-------------")
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")

}
