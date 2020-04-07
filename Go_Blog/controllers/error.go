package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "404.html"
}
