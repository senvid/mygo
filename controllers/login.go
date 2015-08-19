package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	IsExit := c.Input().Get("exit") == "true"
	if IsExit {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplNames = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"
	fmt.Println(username, password, autologin)
	if beego.AppConfig.String("username") ==
		username && beego.AppConfig.String("password") == password {
		maxAge := 0
		if autologin {
			maxAge = 1 << 20
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
		c.Redirect("/", 301)
		return
	}
	// 登陆待定
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {

		return false
	}
	password := ck.Value

	return beego.AppConfig.String("username") ==
		username && beego.AppConfig.String("password") == password
}
