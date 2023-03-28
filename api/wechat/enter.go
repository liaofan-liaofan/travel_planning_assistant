package wechat

import (
	"project/service"
)

type ApiGroup struct {
	WxLoginApi
	ProfilePageApi
	TravelCirclesApi
}

var (
	loginService         = service.ServiceGroupApp.WechatServiceGroup.UserLoginService
	profilePageService   = service.ServiceGroupApp.WechatServiceGroup.ProfileService
	travelCirclesService = service.ServiceGroupApp.WechatServiceGroup.TravelCirclesService
	jwtService           = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
