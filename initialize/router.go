package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"project/global"
	"project/middleware"
	router "project/router"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	Router.StaticFS("/api/"+global.TPA_CONFIG.Local.Path, http.Dir(global.TPA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.TPA_LG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.TPA_LG.Info("use middleware cors")
	Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.TPA_LG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System // 系统后台api
	WxRouter := router.RouterGroupApp.Wechat     // 前台api
	PublicGroup := Router.Group("api")
	{
		systemRouter.InitBaseRouter(PublicGroup)      // 注册基础功能路由 不做鉴权
		WxRouter.InitWxLoginRouter(PublicGroup)       // 微信授权路由
		WxRouter.InitTravelCirclesRouter(PublicGroup) // 旅游圈路由
		//WxRouter.InitProfilePageRouter(PublicGroup) // 个人页面路由注册
	}

	PrivateWxGroup := Router.Group("api")
	PrivateWxGroup.Use(middleware.JWTAuth())
	{
		WxRouter.InitProfilePageRouter(PrivateWxGroup)           // 个人页面路由注册
		WxRouter.InitFileUploadAndDownloadRouter(PrivateWxGroup) // 前台文件上传
	}

	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)               // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)               // 注册menu路由
		systemRouter.InitDeptRouter(PrivateGroup)               //注册部门路由
		systemRouter.InitSystemRouter(PrivateGroup)             // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)             // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)          // 注册角色路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
		systemRouter.InitFileRouter(PrivateGroup)               //文件操作
		//systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
	}

	return Router
}
