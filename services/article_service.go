package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"seagullfly/models"
	"seagullfly/utils"
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
	redis := utils.GetRedisConn()
	redis.Do("SETEX","name",300,"alice")
	return res
}