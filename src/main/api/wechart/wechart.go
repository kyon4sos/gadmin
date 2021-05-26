package wechart

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"neko/src/main/model"
	"neko/src/main/response"
	"neko/src/main/wxmini"
	"strings"
)

const MicroMessenger = "MicroMessenger"

func RegisterWx(r *gin.RouterGroup) {
	r.GET("/login", login)
	r.GET("/access-token", accessToken)
	r.GET("/qrcode", getQrcode)
	r.GET("message-send", subMessageSend)
}

func subMessageSend(ctx *gin.Context) {
	app := wxmini.GetApp("neko")
	k := make(map[string]map[string]interface{})
	k["name1"] = map[string]interface{}{
		"value": "老王",
	}
	k["name3"] = map[string]interface{}{
		"value": "蓝头发",
	}
	k["phone_number4"] = map[string]interface{}{
		"value": 12345678901,
	}
	var req = &wxmini.SubMessageReq{
		Touser:     "oz7Tq0DbhZTOPxLxHJNk0fvPmyek",
		TemplateId: "wFxYkrBOog5ulWTZnHzmpZhed-oHaOckRBVWWTdTBOE",
		Data:       k,
	}
	res, err := app.SubMessageSend(req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response.Ok(ctx, res)
}
func accessToken(ctx *gin.Context) {
	app := wxmini.GetApp("neko")
	token, err := wxmini.GetAccessToken(app)
	if err != nil {
		response.Fail(ctx, "系统错误")
	}
	if token == nil {
		response.Fail(ctx, "token 为空")
	}
	response.Ok(ctx, token)
}
func login(ctx *gin.Context) {
	app := wxmini.GetApp("neko")
	code := ctx.Query("code")
	agent := ctx.Request.UserAgent()
	fmt.Printf("wechart login code %s\n", code)
	res := strings.Contains(agent, MicroMessenger)
	if !res || len(code) == 0 {
		response.Fail(ctx, "请用微信浏览器登录")
		return
	}
	sessionAndOpenId, err := wxmini.GetCode2Session(app,code)
	if err != nil {
		response.FailSyetem(ctx, err)
		return
	}
	var user model.User
	result, findErr := user.FindUserByOpenId(sessionAndOpenId.Openid)
	if findErr != nil {
		response.FailSyetem(ctx, findErr)
		return
	}
	if result == nil {
		user := &model.User{
			SessionKey: sessionAndOpenId.SessionKey,
			Unionid:    sessionAndOpenId.Unionid,
			Openid:     sessionAndOpenId.Openid,
		}
		_, err := user.CreateUser(user)
		if err == nil {
			response.Ok(ctx, user)
			return
		}
		response.FailSyetem(ctx, err)
		return
	}
	response.OkMsg(ctx, "登录成功", user.Username)
}
func getQrcode(ctx *gin.Context) {
	app := wxmini.GetApp("neko")

	qr := &wxmini.QrcodeReq{Scene: "hellowechartapi"}
	res, err := app.GetUnlimitedWxaCode(qr)
	if err != nil {
		response.FailSyetem(ctx, err)
		return
	}
	response.Ok(ctx, res)
}
func handleError(ctx *gin.Context, err error) {
	res := AssertError(err)
	if res!=nil {
		switch res.Errcode {
		case 43101:
			response.Fail(ctx,res.Errmsg)
		}
	}
}
func AssertError(err error) *wxmini.WxRepErr {
	var wxerr=&wxmini.WxRepErr{}
	as := errors.As(err, &wxerr)
	if as {
		fmt.Printf("微信API错误：%v\n",err)
		return wxerr
	}else {
		fmt.Printf("系统错误: %s\n",err)
	}
	return nil
}