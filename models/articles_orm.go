package models

import (
	orm "github.com/astaxie/beego/orm"
	"seagullfly/utils"
	"fmt"
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
	return "article_info"
}

func (this ArticlesOrm) GetQuery() orm.QuerySeter {
	conn := GetConnection()
	return conn.QueryTable(this.TableName())
}

func (this ArticlesOrm) QueryTest() []orm.Params {
	itemList := make([]orm.Params, 0)
	sql := fmt.Sprintf("select * from %s ", this.TableName())

	conn := GetConnection()
	conn.Raw(sql).Values(&itemList)
	return itemList
}