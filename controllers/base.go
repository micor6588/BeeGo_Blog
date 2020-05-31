package controllers

import (
	"github.com/astaxie/beego"
)

type BaseControllers struct {
	beego.Controller
}

func (b *BaseControllers) Prepare() {
	b.Data["Path"] = b.Ctx.Request.RequestURI
}
