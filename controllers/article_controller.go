package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

type ArticleController struct {
	MainController
}

func Init() {
	fmt.Printf("hello beego")
}

func (this ArticleController) List() {
	fmt.Printf("hello beego")
	this.Ctx.WriteString("hello world!!!")
	logs.Info("hello logs")
}