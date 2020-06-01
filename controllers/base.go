package controllers

import (
	"errors"
	"liteblog/syserrors"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

type BaseControllers struct {
	beego.Controller
}

func (b *BaseControllers) Prepare() {
	b.Data["Path"] = b.Ctx.Request.RequestURI
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (c *BaseControllers) GetMustString(key string, msg string) string {
	email := c.GetString(key, "")
	if len(email) == 0 {
		c.Abort500(errors.New(msg))
	}
	return email
}

func (ctx *BaseControllers) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

func (ctx *BaseControllers) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (ctx *BaseControllers) JSONOkH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	ctx.Data["json"] = maps
	ctx.ServeJSON()
}

func (ctx *BaseControllers) JSONOkData(count int, data interface{}) {
	ctx.Data["json"] = &ResultJsonValue{
		Code:  0,
		Count: count,
		Msg:   "成功！",
		Data:  data,
	}
	ctx.ServeJSON()
}

func (this *BaseControllers) UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		this.Abort500(syserrors.NewError("系统错误", err))
	}
	return u.String()
}
