package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplNames = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
