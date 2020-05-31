package controllers

type IndexController struct {
	BaseControllers
}

// @router / [get]
func (c *IndexController) Index() {
	c.TplName = "index.html"
}

// @router /message [get]
func (c *IndexController) IndeMessage() {

	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) About() {

	c.TplName = "about.html"
}
