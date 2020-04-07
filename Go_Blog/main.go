package main

import (
	_ "BeeGO_Blog/Go_Blog/models"
	_ "BeeGO_Blog/Go_Blog/routers"
	"strings"

	"github.com/astaxie/beego"
)

func main() {
	initTemplate()
	beego.Run()
}

//比较两个路径是否相等
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}
