package router

import (
	"project/router/system"
	"project/router/wechat"
)

type RouterGroup struct {
	System    system.RouterGroup
	FrontDesk wechat.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
