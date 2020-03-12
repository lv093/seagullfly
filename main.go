package main

import (
	_ "seagullfly/routers"
	"github.com/astaxie/beego"

	"seagullfly/initialize"
	"net/http"
)

func init() {
	initialize.RegisterDatabase()
	initialize.RegisterRedis()

	//beego.SetStaticPath("/uploads", "uploads")
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		go func() {
			http.ListenAndServe("localhost:8088", nil)
		}()
	}

	beego.BConfig.Listen.Graceful = false
	beego.Run()
}

