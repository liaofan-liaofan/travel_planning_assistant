package wechat

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type UserLogin struct{}

func (u *UserLogin) WxLogin(Router *gin.RouterGroup) {
	baseRouter := Router.Group("wx")
	var baseApi = v1.ApiGroupApp.Wechat.UserLogin
	{
		baseRouter.POST("login", baseApi.Login) // 微信授权登录
	}
}
