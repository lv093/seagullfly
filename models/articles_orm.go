package models

import (
	orm "github.com/astaxie/beego/orm"
	"seagullfly/utils"
)

func init() {
	orm.RegisterModel(new(ArticlesOrm))
}

func GetConnection() orm.Ormer {
	conn := utils.GetDbConnection("default")
	return conn
}

type ArticlesOrm struct {
	Id        int
	SdId      int
	Code      int
	Name      string
	Logo      string
	CreatedAt string
	UpdatedAt string
}

func (this ArticlesOrm) TableName() string {
	return "cfa_teams"
}

func (this ArticlesOrm) GetQuery() orm.QuerySeter {
	conn := GetConnection()
	return conn.QueryTable(this.TableName())
}
