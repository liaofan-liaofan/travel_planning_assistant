package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	systemRes "project/model/system/response"
	"project/utils"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {string}  string  "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router    /api/base/captcha [post]

func (b *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.TPA_CONFIG.Captcha.ImgHeight, global.TPA_CONFIG.Captcha.ImgWidth, global.TPA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	global.TPA_LG.Info("requestId获取成功", utils.GetRequestID(c))
	if id, b64s, err := cp.Generate(); err != nil {
		global.TPA_LG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
