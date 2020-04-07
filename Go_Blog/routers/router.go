package routers

import (
	"BeeGO_Blog/Go_Blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
}
