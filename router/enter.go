package router

import (
	"project/router/system"
	"project/router/wechat"
)

type RouterGroup struct {
	System system.RouterGroup
	Wechat wechat.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
