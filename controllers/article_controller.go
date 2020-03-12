package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"seagullfly/services"
)

type ArticleController struct {
	BaseController
}

func Init() {
	fmt.Printf("hello beego")
}

func (this ArticleController) List() {
	articleService := services.GetArticleService()
	articles := articleService.TestConds()
	logs.Info("articles ret:%v", articles)
	this.Ctx.Output.JSON(articles, true,true)
}