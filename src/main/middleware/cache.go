package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Cache(key interface{}) gin.HandlerFunc {
	fmt.Printf("%s 缓存中间件 \n",time.Now())
	return func(ctx *gin.Context) {
		fmt.Printf("缓存中间件 \n")
		switch ctx.Request.Method {
		case "POST":
			cachePost(ctx)
		case "GET":
			cacheGet(ctx)
		case "DEL":
			cacheEvict(ctx)
		case "PUT":
			cachePut(ctx)
		default:
			ctx.Next()
		}
	}
}

func cacheEvict(ctx *gin.Context) {

}

func cacheGet(ctx *gin.Context) {

}

func cachePost(ctx *gin.Context) {

}
func cachePut(ctx *gin.Context) {

}
