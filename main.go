package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistrtDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
