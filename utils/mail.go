package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-gomail/gomail"
	"fmt"
)

var style = "<style> table { width: 80%; font-size: .938em; text-align: center; border-collapse: collapse; } caption { margin: 1em 0 .7em 0; text-align: center; font-weight: bold; font-size: 120%; letter-spacing: .5px; color: #fff; } th { text-align: center; padding: .5em .5em; font-weight: bold; background: #808080; color: #fff; } td { padding: .5em .5em; border-bottom: solid 1px #000; } table, table tr th, table tr td { border: 1px solid #d0d0d0; } </style>"

func SendTableMail(subject string, mailContent string) string {
	var toEmails []string
	content := "<html>" + style + "<table>" + mailContent + "</table></html>"
	if beego.BConfig.RunMode == "beta" {
		toEmails = []string{"lottery@dongqiudi.com", "sport-data@dongqiudi.com"}
	} else {
		toEmails = []string{"lulu@dongqiudi.com", "liuyunzhi@dongqiudi.com"}
	}
	err := SendMail(toEmails, subject, content, "html")
	fmt.Println(err)
	if err != nil {
		logs.Info("Send HtmlMail err", err)
	} else {
		logs.Info("Send HtmlMail ok", toEmails)
	}
	return content
}

func SendMail(to []string, subject, body, mailType string) error {
	user := "736598912@qq.com"
	password := "xxx"
	//host := "smtp.exmail.qq.com"
	host := "smtp.qq.com"
	port := 465

	contentType := "text/plain"
	if mailType == "html" {
		contentType = "text/html"
	}

	m := gomail.NewMessage()

	m.SetHeader("From", user)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody(contentType, body)
	d := gomail.NewDialer(host, port, user, password)

	err := d.DialAndSend(m)
	fmt.Println(err)
	return err
}
