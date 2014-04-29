package main

import (
	"github.com/astaxie/beego"
	_ "weixin/routers"
)

func main() {
	beego.Trace("Log is ok")
	beego.Run()
}
