package routers

import (
	"github.com/astaxie/beego"
	"weixin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin", &controllers.WeixinController{})
	beego.Router("/weixin/getTextMessage", &controllers.TextMessageController{})
}
