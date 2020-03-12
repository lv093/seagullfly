package controllers

import (
	"github.com/astaxie/beego/logs"
	"crypto/sha1"
	"fmt"
)

type AccountController struct {
	BaseController
}


func (this AccountController) WechatVerify() {
	params := make(map[string]string, 0)
	params["nonce"] = this.Ctx.Input.Query("nonce")
	params["timestamp"] = this.Ctx.Input.Query("timestamp")
	params["echostr"] = this.Ctx.Input.Query("echostr")
	params["signature"] = this.Ctx.Input.Query("signature")

	//校验参数
	calSig := sha1.New()
	calSig.Write([]byte(params["nonce"] + params["timestamp"] + "seagull2020"))
	sigRes := calSig.Sum(nil)
	sigStr := fmt.Sprintf("%x\n", sigRes)
	logs.Info("wechat verify info:", sigStr, params)
	if sigStr == params["signature"] && params["echostr"] != "" {
		this.Ctx.Output.Body([]byte(params["echostr"]))
	}
	this.Ctx.Output.Body([]byte("failed"))
}