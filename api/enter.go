package api

import "project/api/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
