package initialize

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func RegisterLog() {
	path := beego.AppConfig.String("log.path")
	cfg := `{"filename":"` + path + `","separate":["critical", "error", "debug"],"perm":"644"}`

	err := logs.SetLogger(logs.AdapterMultiFile, cfg)
	logs.EnableFuncCallDepth(true)
	beego.Info("Lib:Log Init, " + path + ", Start!")
	if err != nil {
		beego.Error("Lib:Log Init, " + path + ", Failed!")
	}
	beego.Info("Lib:Log Init, " + path + ", Finish!")
	if beego.BConfig.RunMode != "dev" {
		beego.BeeLogger.DelLogger("console")
	}
	beego.BConfig.Log.AccessLogs = true
}
