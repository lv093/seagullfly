package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"seagullfly/models"
)

var articleService *ArticleService

func GetArticleService() *ArticleService {
	if articleService == nil {
		logs.Info("init ArticleService")
		articleService = new(ArticleService)
	}
	return articleService
}

type ArticleService struct {
}

func (this ArticleService) List() []orm.Params {
	articleOrm := new(models.ArticlesOrm)
	res := articleOrm.QueryTest()
	return res
}