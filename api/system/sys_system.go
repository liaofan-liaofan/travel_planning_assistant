package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/model/system"
	systemRes "project/model/system/response"
	"project/utils"
)

type SystemApi struct{}

// GetSystemConfig @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// SetSystemConfig @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /api/system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	_ = c.ShouldBindJSON(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.TPA_LG.Error("设置失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithData("设置成功", c)
	}
}

// ReloadSystem @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"重启系统成功"}"
// @Router /api/system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.TPA_LG.Error("重启系统失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("重启系统失败", c)
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
}

// GetServerInfo @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
