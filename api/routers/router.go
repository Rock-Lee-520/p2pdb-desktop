package routers

import (
	"api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "*:Index")
	beego.Router("/email", &controllers.MainController{}, "*:Email")
	beego.Router("/compose", &controllers.MainController{}, "*:Compose")
	beego.Router("/chat", &controllers.MainController{}, "*:Chat")
}
