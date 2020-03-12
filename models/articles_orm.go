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

/**
CREATE TABLE `article_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'article id',
  `topic_id` int(10) NOT NULL DEFAULT '0' COMMENT '对应专题',
  `author` varchar(50) NOT NULL DEFAULT '',
  `title` varchar(500) NOT NULL DEFAULT '',
  `content` text,
  `status` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `published_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1199 DEFAULT CHARSET=utf8;
 */
type ArticlesOrm struct {
	Id        	int
	TopicId     int
	Author      string
	Title      	string
	Content     string
	Status      string
	CreatedAt 	string
	UpdatedAt 	string
	PublishedAt string
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