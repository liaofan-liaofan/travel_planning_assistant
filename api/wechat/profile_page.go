package wechat

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/utils"
)

type ProfilePageApi struct{}

// @Tags ProfilePageApi
// @Summary 微信小程序登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/profile-page/ [get]

func (p *ProfilePageApi) GetProfileApi(c *gin.Context) {
	userId := utils.GetWxUserID(c)
	if userId == 0 {
		response.FailWithMessage("中间件获取用户ID错误", c)
	}
	// 业务逻辑处理
	if result, err := profilePageService.SelectProfileByID(userId); err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any(
			"err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}
