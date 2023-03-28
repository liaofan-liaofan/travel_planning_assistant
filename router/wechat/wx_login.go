package wechat

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type WxLoginRouter struct{}

func (u *WxLoginRouter) InitWxLoginRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("wx")
	var baseApi = v1.ApiGroupApp.Wechat.WxLoginApi
	{
		baseRouter.POST("login", baseApi.Login) // 微信授权登录
	}
}
