package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeeGO_Blog/Go_Blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeeGO_Blog/Go_Blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGO_Blog/Go_Blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["BeeGO_Blog/Go_Blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
