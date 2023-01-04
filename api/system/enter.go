package system

import "project/service"

type ApiGroup struct {
	SystemApiApi
	BaseApi
}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
