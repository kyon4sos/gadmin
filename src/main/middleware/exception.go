package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"neko/src/main/exception"
	"time"
)

func Exception() gin.HandlerFunc {
	fmt.Printf("%s exception handler \n",time.Now())
	return func(ctx *gin.Context) {
		if err:=recover();err!=nil {
			switch err.(type) {
			case exception.ErrParam:
				ctx.JSON(200,gin.H{"msg":"此参数异常"})
			}
		}
	}
}