package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"neko/src/main/api"
	"neko/src/main/exception"
	"neko/src/main/wxmini"
	"time"
)
func main() {
	app, _ :=wxmini.NewApp("wxc771362245f54f98",
		"0489a75011688b73f730f7f38e968244")
	app.Register("neko")
	gin.ForceConsoleColor()
	http := gin.Default()
	cache := make(map[string]string)
	cache["user"] = "login"
	//http.Use(middleware.Cache(cache))
	http.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	http.Use(gin.Recovery())
	http.NoMethod(handlerNoFound)
	http.NoRoute(handlerNoFound)

	//http.Use(TlsHandler())
	//go https.RunTLS(":9090", "config/cert.pem", "config/cert.key")
	//api.InitController(http)
	//api.InitWxController(http)
	api.InitRouter(http)
	//http.Run(":8080")
	http.RunTLS(":99","config/cert.pem", "config/cert.key")

}

func handlerNoFound(ctx *gin.Context) {
	res := ctx.Request.Method + " " + ctx.Request.URL.String()
	ctx.JSON(200, gin.H{
		"msg":  "走错地方了",
		"data": res,
	})
	return
}

//
func exceptionHandler(ctx *gin.Context) {
	fmt.Printf("异常中间件 \n")
	if err := recover(); err != nil {
		switch err.(type) {
		case exception.ErrParam:
			ctx.JSON(200, gin.H{"msg": "此参数异常"})
		}
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:99",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}