package main

import (
	_ "seagullfly/routers"
	"github.com/astaxie/beego"

	"seagullfly/initialize"
	"net/http"
)

func init() {
	initialize.RegisterDatabase()
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		go func() {
			http.ListenAndServe("192.168.0.112:8031", nil)
		}()
	}

	beego.BConfig.Listen.Graceful = false
	beego.Run()
}

