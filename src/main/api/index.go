package api

import (
	"github.com/gin-gonic/gin"
	"neko/src/main/cache"
	"neko/src/main/response"
)

func registerIndex(group *gin.RouterGroup) {
	group.GET("/admin", login)
}
func login(ctx *gin.Context) {
	query := ctx.Query("name")
	get, _ := cache.Get(query)
	if get != nil {
		response.Ok(ctx, get)
		return
	}
	//user := model.User{UserName: query}
	////err := user.FindByUsername()
	////if err !=nil {
	////	response.Fail(ctx,"用户名或者密码错误")
	////	return
	////}
	//cache.SetPreFixKey(user.UserName, &user)
	//prefix := user.GetCachePrefix()
	//fmt.Printf("prefix %s \n",prefix)
	//user.SetCachePrefix("这是有一个")
	//prefix = user.GetCachePrefix()
	//fmt.Printf("prefix %s \n",prefix)
	//response.Ok(ctx,&user)
	return
}
