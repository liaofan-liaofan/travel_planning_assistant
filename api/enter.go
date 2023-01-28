package api

import (
	"project/api/system"
	"project/api/wechat"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	Wechat         wechat.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
