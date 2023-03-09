package wechat

import "project/service"

type ApiGroup struct {
	UserLogin
}

var loginService = service.ServiceGroupApp.WechatServiceGroup.UserLogin
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
