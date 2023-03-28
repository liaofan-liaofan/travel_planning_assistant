package wechat

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type ProfilePageRouter struct{}

func (p *ProfilePageRouter) InitProfilePageRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("profile-page")
	var baseApi = v1.ApiGroupApp.Wechat.ProfilePageApi
	{
		baseRouter.GET("/", baseApi.GetProfileApi) // 微信授权登录
	}
}
