package wechat

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
	"project/middleware"
)

type TravelCirclesRouter struct{}

func (t *TravelCirclesRouter) InitTravelCirclesRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("travel-circles")
	var baseApi = v1.ApiGroupApp.Wechat.TravelCirclesApi
	{
		baseRouter.GET("travel-post", baseApi.SearchTravelCircles)                // 查询旅游帖(游客)api
		baseRouter.GET("travel-post-details", baseApi.SearchTravelCirclesDetails) // 查询旅游圈详情(游客)api
	}
	baseRouter.Use(middleware.JWTAuth())
	{
		baseRouter.GET("travel-postJwt", baseApi.SearchTravelCirclesJwt)   // 查询旅游圈(登录)api
		baseRouter.GET("single-playerShow", baseApi.SearchPlayerShowById)  // 根据旅游圈ID查询玩家秀
		baseRouter.GET("detailsAu", baseApi.SearchPlayerShowDetailsByIdAu) // 查询旅游圈详情(登录)api
		baseRouter.PUT("details", baseApi.LikePlayerShowById)              // 给旅游圈点赞api
		baseRouter.POST("new", baseApi.NewPlayerShow)                      // 确认发布旅游帖
	}
}
