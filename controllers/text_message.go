package controllers

import (
	//"code.google.com/p/go.text/encoding"
	//"code.google.com/p/go.text/encoding/charmap"
	//"code.google.com/p/go.text/transform"
	"github.com/astaxie/beego"
	iconv "github.com/djimenez/iconv-go"
	"strings"
)

type TextMessageController struct {
	beego.Controller
}

func (this *TextMessageController) Post() {
	r := this.Ctx.Request
	//this.ParseForm()
	r.ParseForm()
	//beego.Trace("ToUserName", this.GetString("ToUserName"))
	//beego.Trace("ToUserName", r.PostFormValue("ToUserName"))
	//beego.Trace("MsgType", this.GetString("MsgType"))
	//beego.Trace("MsgType", r.PostFormValue("MsgType"))
	for k, v := range r.Form {
		beego.Trace("key:", k)
		beego.Trace("val:", strings.Join(v, ""))
	}
	//for k, v := range r.PostForm {
	//	beego.Trace("key:", k)
	//	beego.Trace("val:", strings.Join(v, ""))
	//}
	//for i, s := range string(this.Ctx.Input.RequestBody) {
	//	beego.Trace(i, ":", s, ":", string(s))
	//}
	//beego.Trace("body:", string(this.Ctx.Input.RequestBody))
	output, _ := iconv.ConvertString(string(this.Ctx.Input.RequestBody), "utf8", "gbk")
	//beego.Trace(charmap.Windows1252.NewEncoder())
	//encoding.UTF8Validator.Transform()
	beego.Trace(output)
	this.Data["Out"] = "sdfasdfasdfasdfasdfa"
	this.TplNames = "out.tpl"
}

func (this *TextMessageController) Get() {
	this.Post()
}
