package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type TextMessageController struct {
	beego.Controller
}

func (this *TextMessageController) Post() {
	r := this.Ctx.Request
	//this.ParseForm()
	r.ParseForm()
	beego.Trace("ToUserName", this.GetString("ToUserName"))
	beego.Trace("ToUserName", r.PostFormValue("ToUserName"))
	for k, v := range r.PostForm {
		beego.Trace("key:", k)
		beego.Trace("val:", strings.Join(v, ""))
	}
	this.Data["Out"] = "sdfasdfasdfasdfasdfa"
	this.TplNames = "out.tpl"
}

func (this *TextMessageController) Get() {
	this.Post()
}
