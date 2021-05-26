package api

import (
	"github.com/gin-gonic/gin"
	"neko/src/main/api/admin"
	"neko/src/main/api/wechart"
)

func InitRouter(engine *gin.Engine)  {
	adminGroup:=engine.Group("/admin")
	wxGroup :=engine.Group("/wechart")
	wechart.RegisterWx(wxGroup)
	admin.RegisterAdmin(adminGroup)
}