package routers

import (
	"weizhi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/return", &controllers.MainController{},"post:GetUrl")
}
