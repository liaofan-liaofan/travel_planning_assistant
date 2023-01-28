package wechat

import (
	"github.com/gin-gonic/gin"
	"project/model/wechat/request"
	"project/model/wechat/response"
)

type UserLogin struct{}

// WXLogin 微信登陆
func (w *UserLogin) WXLogin(c *gin.Context, user *request.WxCode, wxLoginInfo *response.WXLoginResp) (middleInfo response.UserInfo, err error) {
	//middleInfo, err = WxUserService.InsertWxUser(wxLoginInfo, user)
	//if err != nil {
	//	zap.L().Error("wx InsertWxUser failed")
	//	app.ResponseError(c, app.CodeInsertOperationFail)
	//	return middleInfo, err
	//}
	return middleInfo, err
}
