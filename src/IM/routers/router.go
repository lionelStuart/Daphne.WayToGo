package routers

import (
	"IM/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/join", &controllers.MainController{}, "post:Join")

	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
