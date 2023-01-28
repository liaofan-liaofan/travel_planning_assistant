package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"project/global"
	cres "project/model/common/response"
	"project/model/wechat/request"
	wres "project/model/wechat/response"
	"project/utils"
)

type UserLogin struct{}

// @Tags FrontDeskApi
// @Summary 微信小程序登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/wx/login [post]

func (u *UserLogin) Login(c *gin.Context) {
	var data request.WxCode
	_ = c.ShouldBindJSON(&data)
	if err := utils.Verify(data, utils.IdVerify); err != nil {
		cres.FailWithMessage(err.Error(), c)
		return
	}

	// 根据code获取 openID 和 session_key
	wxLoginInfo, err := HttpGetRequest(&data)
	if err != nil || wxLoginInfo.OpenId == "" {
		cres.FailWithMessage(err.Error(), c)
		return
	}

	// 业务逻辑处理
	result, err := loginService.WXLogin(c, &data, &wxLoginInfo)
	if err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any("err", err), utils.GetRequestID(c))
		cres.FailWithMessage("获取失败", c)
	} else {
		cres.OkWithData(result, c)
	}
}

// HttpGetRequest HTTP GET请求带参数
func HttpGetRequest(wx *request.WxCode) (wxLoginInfo wres.WXLoginResp, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, global.TPA_CONFIG.WxLogin.AppId, global.TPA_CONFIG.WxLogin.Secret, wx.JsCode)
	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return wres.WXLoginResp{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxLoginInfo); err != nil {
		log.Println(err)
		return wxLoginInfo, err
	}
	return wxLoginInfo, nil
}
