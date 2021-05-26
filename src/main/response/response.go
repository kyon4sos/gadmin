package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Resoponse struct {
	Code int
	Msg  string
}


func OkMsg(ctx *gin.Context,msg string,data interface{}) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  MsgOk,
		"code":CodeOk,
		"data":data,
	})
}

func Ok(ctx *gin.Context,data interface{}) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  MsgOk,
		"code":CodeOk,
		"data":data,
	})
	return
}
func OkAndMsg(ctx *gin.Context,msg string,data interface{}) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  msg,
		"code":CodeOk,
		"data":data,
	})
	return
}
func OkComplete(ctx *gin.Context,msg string,code int,data interface{}) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  msg,
		"code":code,
		"data":data,
	})
	return
}
func OkAndCode(ctx *gin.Context,code int,data interface{}) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  MsgOk,
		"code":code,
		"data":data,
	})
	return
}
func FailSyetem(ctx *gin.Context,err error) {
	fmt.Printf("系统错误：%s",err.Error())
	ctx.JSON(StatusOk, gin.H{
		"msg":  "系统错误",
		"code":CodeFail,
	})
	return
}
func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(StatusOk, gin.H{
		"msg":  msg,
		"code":CodeFail,
	})
	return
}
func FailMsg(ctx *gin.Context, code int, msg string, ) {
	ctx.JSON(StatusFail, gin.H{
		"code": code,
		"msg":  msg,
	})
	return
}

const (
	StatusOk   = 200
	StatusFail = 200
)
const (
	CodeOk = 2000
	CodeFail =4000
)
const (
	MsgFail = "操作失败"
	MsgOk   = "操作成功"
)
