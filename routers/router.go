package routers

import (
	"seagullfly/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/article/list", &controllers.ArticleController{},"Get:List")


    beego.Router("/account/wechat/verify", &controllers.AccountController{},"Get:WechatVerify")

}
