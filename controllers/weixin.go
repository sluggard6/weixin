package controllers

import (
	"github.com/astaxie/beego"
)

type WeixinController struct {
	beego.Controller
}

func (this *WeixinController) Get() {
	this.Data["Out"] = "My Out ................................."
	this.TplNames = "out.tpl"
}
