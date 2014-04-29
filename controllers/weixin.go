package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

const TOKEN = "guanglanludian"

type WeixinController struct {
	beego.Controller
}

func (this *WeixinController) Get() {
	r := this.Ctx.Request
	r.ParseForm()
	for k, v := range r.Form {
		beego.Trace("key:", k)
		beego.Trace("val:", strings.Join(v, ""))
	}
	timestamp, nonce, echostr, signature := this.GetString("timestamp"), this.GetString("nonce"), this.GetString("echostr"), this.GetString("signature")
	slice := []string{TOKEN, timestamp, nonce}
	beego.Trace("before sort:", slice)
	sort.Strings(slice)
	//sort.Sort([]byte(slice))
	beego.Trace("after sort:", slice)
	tempStr := slice[0] + slice[1] + slice[2]
	beego.Trace("tempStr:", tempStr)
	t := sha1.New()
	io.WriteString(t, tempStr)
	tempStr = fmt.Sprintf("%x", t.Sum(nil))
	this.Data["Out"] = "error"
	if strings.EqualFold(tempStr, signature) {
		this.Data["Out"] = echostr
	}
	this.TplNames = "out.tpl"
}
