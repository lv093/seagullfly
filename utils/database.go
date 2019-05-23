package utils

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

func GetDbConnection(name string) orm.Ormer {
	orm := orm.NewOrm()
	// TODO 判断是否在在注册的列表里, 有Error的话, 需要进行错误处理
	err := orm.Using(name)
	if err != nil {
		logs.Error("GetDbConnection, Error, ", name, err)
	}
	return orm
}
