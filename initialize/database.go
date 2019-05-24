package initialize

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDatabase() {
	logs.Info("Lib:Database Init, Start!")
	initMysql()
	logs.Info("Lib:Database Init, Finish!")
}

func initMysql() bool {
	logs.Info("Lib:Database MySQL Init, Start!")

	dbs := GetDbList()
	for k := range dbs {
		v := dbs[k]

		if !initMysqlDatabase(v) {
			return false
		}
	}

	logs.Info("Lib:Database MySQL Init, Finish!")

	runMode := beego.AppConfig.String("RunMode")
	if runMode == "test" || runMode == "dev" {
		orm.Debug = true
	}
	return true
}

func initMysqlDatabase(name string) bool {
	logs.Info("Lib:Database MySQL Init, " + name + ", Start!")

	dbDriver := beego.AppConfig.String("db." + name + ".driver")
	dbHost := beego.AppConfig.String("db." + name + ".host")
	dbPort := beego.AppConfig.String("db." + name + ".port")
	dbUser := beego.AppConfig.String("db." + name + ".user")
	dbPwd := beego.AppConfig.String("db." + name + ".pwd")
	dbName := beego.AppConfig.String("db." + name + ".name")
	dbMaxIdle, _ := beego.AppConfig.Int("db." + name + ".max_idle")
	dbMaxConn, _ := beego.AppConfig.Int("db." + name + ".max_conn")

	link := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase(name, dbDriver, link, dbMaxIdle, dbMaxConn)

	if err != nil {
		logs.Error("Lib:Database MySQL Init, " + name + ", Failed!")
		return false
	}

	logs.Info("Lib:Database MySQL Init, " + name + ", Finish!")
	return true
}

func GetDbList() []string {
	return []string{"default"}
}
