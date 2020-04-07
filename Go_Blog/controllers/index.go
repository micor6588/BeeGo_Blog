package controllers

type IndexController struct {
	BaseController
}

// @router /about [get]
func (c *IndexController) IndexAbout() {
	c.TplName = "about.html"
}

// @router /message [get]
func (c *IndexController) IndexMessage() {
	c.TplName = "message.html"
}

// @router / [get]
func (c *IndexController) Index() {
	c.TplName = "index.html"
}
