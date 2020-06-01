package controllers

type IndexController struct {
	BaseControllers
}

// @router / [get]
func (c *IndexController) Index() {

	c.TplName = "index.html"
}

// @router /message [get]
func (c *IndexController) IndexMessage() {
	c.Abort("404")
	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) About() {
	c.Abort("500")
	c.TplName = "about.html"
}
