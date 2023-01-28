package service

import (
	"project/service/system"
	"project/service/wechat"
)

type ServiceGroup struct {
	SystemServiceGroup system.SysGroup
	WechatServiceGroup wechat.WechatGroup
}

var ServiceGroupApp = new(ServiceGroup)
