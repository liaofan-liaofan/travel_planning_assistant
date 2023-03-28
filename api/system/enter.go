package system

import (
	"project/service"
	"project/utils"
)

type ApiGroup struct {
	BaseApi
	JwtApi
	SystemApiApi
	AuthorityMenuApi
	DeptApi
	SystemApi
	CasbinApi
	AuthorityApi
	OperationRecordApi
	SysFileUploadAndDownloadApi
}

var dataScope = utils.DataScope{}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var DeptService = service.ServiceGroupApp.SystemServiceGroup.DeptService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var fileUploadAndDownloadService = service.ServiceGroupApp.SystemServiceGroup.FileUploadAndDownloadService
