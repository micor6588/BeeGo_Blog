package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

// @router /about [get]
func (c *IndexController) IndexAbout() {
	c.TplName = "about.html"
}

// @router / [get]
func (c *IndexController) Index() {
	c.TplName = "index.html"
}
