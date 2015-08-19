package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		category := c.Input().Get("category")
		if len(category) == 0 {
			break
		}
		err := models.AddCategory(category)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return

	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
	}

	c.Data["IsCategory"] = true
	c.TplNames = "category.html"
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
