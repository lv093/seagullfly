package main

import (
	_ "seagullfly/routers"
	"github.com/astaxie/beego"

	"seagullfly/initialize"
)

func init() {
	initialize.RegisterDatabase()
}
func main() {
	beego.Run()
}

