package wechat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/model/wechat/request"
	"project/utils"
)

type TravelCirclesApi struct{}

// @Tags TravelCirclesApi
// @Summary 获取旅游帖（游客）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body result []response.TravelPost, err error true "获取旅游帖信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/travel-circles/travel-post [get]

func (a *TravelCirclesApi) SearchTravelCircles(c *gin.Context) {
	var data request.SearchType
	_ = c.ShouldBindQuery(&data)
	if err := utils.Verify(data, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := travelCirclesService.SearchTravelPost(data)
	if err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any(
			"err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}

// @Tags TravelCirclesApi
// @Summary 获取玩家秀详情（游客）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body result []response.TravelPost, err error true "获取玩家秀详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/travel-circles/travel-post-details [get]

func (a *TravelCirclesApi) SearchTravelCirclesDetails(c *gin.Context) {

}

// @Tags TravelCirclesApi
// @Summary 获取旅游帖（游客）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body result []response.TravelPost, err error true "获取旅游帖信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/travel-circles/travel-postJwt [get]

func (a *TravelCirclesApi) SearchTravelCirclesJwt(c *gin.Context) {
	var data request.SearchType
	userId := utils.GetWxUserID(c)
	if userId == 0 {
		response.FailWithMessage("中间件获取用户ID错误", c)
	}
	_ = c.ShouldBindQuery(&data)
	if err := utils.Verify(data, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("拿到数据", data)
	result, err := travelCirclesService.SearchTravelPostAndLikeStatus(userId, data)
	if err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any(
			"err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}
}

func (a *TravelCirclesApi) SearchPlayerShowById(c *gin.Context) {

}

func (a *TravelCirclesApi) SearchPlayerShowDetailsByIdAu(c *gin.Context) {

}

func (a *TravelCirclesApi) LikePlayerShowById(c *gin.Context) {

}

func (a *TravelCirclesApi) NewPlayerShow(c *gin.Context) {
	var data request.NewPlayShow
	userId := utils.GetWxUserID(c)
	if userId == 0 {
		response.FailWithMessage("中间件获取用户ID错误", c)
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage("数据获取错误", c)
	}
	fmt.Println(data)
	if err := utils.Verify(data, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := travelCirclesService.ConfirmNewPlayerShow(userId, data)
	if err != nil {
		global.TPA_LG.Error("上传失败!", zap.Any(
			"err", err), utils.GetRequestID(c))
		response.FailWithMessage("上传失败", c)
	} else {
		response.OkWithMessage("上传成功", c)
	}
}
