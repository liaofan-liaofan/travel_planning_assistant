package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"project/global"
	"project/middleware"
	routerd "project/router"
)

func Routers() *gin.Engine {
	router := gin.New()
	router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	router.StaticFS("/api/"+global.TPA_CONFIG.Local.Path, http.Dir(global.TPA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.TPA_LG.Info("use middleware logger")
	// 跨域
	router.Use(middleware.Cors()) // 如需跨域可以打开
	global.TPA_LG.Info("use middleware cors")
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.TPA_LG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := routerd.RouterGroupApp.System
	PublicGroup := router.Group("api")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup) // 注册功能api路由
	}
	return router
}
