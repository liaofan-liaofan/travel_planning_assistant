package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/model/system"
	"project/utils"
)

type SystemApiApi struct{}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/api/createApi [post]

func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		global.TPA_LG.Error("创建失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
