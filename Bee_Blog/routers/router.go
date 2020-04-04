package routers

import (
	"BeeGO_Blog/Bee_Blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
