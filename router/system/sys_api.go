package system

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
	"project/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	var apiRouterApi = v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi) // 创建Api
	}
}
