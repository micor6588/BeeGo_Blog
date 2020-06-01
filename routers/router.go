package routers

import (
	"liteblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.ErrorController{})
}
