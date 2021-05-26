package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"neko/src/main/response"
)

func RegisterAdmin(router *gin.RouterGroup) {
	router.GET("/test", testWX)
}

func testWX(ctx *gin.Context) {
	fmt.Printf("wxin test\n")
	response.OkMsg(ctx,"测试成功",nil)
}