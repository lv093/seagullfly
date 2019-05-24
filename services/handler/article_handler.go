package handler

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"seagullfly/models"
)

var articleDataHandler *ArticleDataHandler

type ArticleDataHandler struct {
}

func GetArticleDataHandler() *ArticleDataHandler {
	if articleDataHandler == nil {
		logs.Info("init handler.ArticleDataHandler")
		articleDataHandler = new(ArticleDataHandler)
	}
	return articleDataHandler
}

func (this *ArticleDataHandler) Handle(msg *models.ArticlesOrm) error {

	if msg == nil {
		return errors.New("message is nil")
	}

	logs.Info("GetArticleDataHandler, Handle, start, ", msg)

	var err error
	switch msg.Name {
	case "update_team":
	}

	logs.Info("GetArticleDataHandler, Handle, finish, ", msg)
	return err
}
