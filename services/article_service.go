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

func (this *ArticleService) TestConds() []*models.ArticlesOrm {
	articles := make([]*models.ArticlesOrm, 0)
	cond := new(orm.Condition)
	cond1 := cond.Or("Status", 0).Or("Status", 1)
	//cond2 := cond.And("TopicId", 2).And("Author", "lv")
	cond3 := cond.AndCond(cond1).And("Author", "lv")
	orm.NewOrm().QueryTable("article_info").SetCond(cond3).OrderBy("-CreatedAt").All(&articles)
	return articles
}