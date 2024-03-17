package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"time1043/gvb-server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) // swag

	apiRouterGroup := router.Group("api")         // 路由分组
	routerGroupApp := RouterGroup{apiRouterGroup} // 路由分层
	routerGroupApp.SettingsRouter()               // 系统配置api
	routerGroupApp.ImagesRouter()                 // 图片管理api
	routerGroupApp.AdvertRouter()                 // 广告管理api
	routerGroupApp.MenuRouter()                   // 菜单管理api

	return router
}
