package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"project/global"
	"project/middleware"
	cres "project/model/common/response"
	"project/model/system"
	systemReq "project/model/system/request"
	"project/model/wechat"
	"project/model/wechat/request"
	wres "project/model/wechat/response"
	"project/utils"
	"time"
)

type WxLoginApi struct{}

// @Tags UserLoginApi
// @Summary 微信小程序登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/wx/login [post]

func (u *WxLoginApi) Login(c *gin.Context) {
	var data request.WxCode
	_ = c.ShouldBindJSON(&data)
	if err := utils.Verify(data, utils.IdVerify); err != nil {
		cres.FailWithMessage(err.Error(), c)
		return
	}
	// 根据code获取 openID 和 session_key
	wxLoginInfo, err := u.HttpGetRequest(&data)
	if err != nil || wxLoginInfo.OpenId == "" {
		cres.FailWithMessage(err.Error(), c)
		return
	}

	// 业务逻辑处理
	if user, err := loginService.WXLogin(data, &wxLoginInfo); err != nil {
		global.TPA_LG.Error("获取失败!", zap.Any(
			"err", err), utils.GetRequestID(c))
		cres.FailWithMessage("获取失败", c)
	} else {
		u.tokenNext(c, user) // 生成token
	}
}

// 登录以后签发jwt
func (b *WxLoginApi) tokenNext(c *gin.Context, user wechat.WxUserInfo) {
	j := &middleware.JWT{SigningKey: []byte(global.TPA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := systemReq.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Username:   user.NickName,
		BufferTime: global.TPA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.TPA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "gsdPlus",                                             // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.TPA_LG.Error("获取token失败!", zap.Any("err", err), utils.GetRequestID(c))
		cres.FailWithMessage("获取token失败", c)
		return
	}
	userCache := systemReq.UserCache{
		ID:   user.ID,
		UUID: user.UUID.String(),
	}
	_ = jwtService.SetRedisUserInfo(user.UUID.String(), userCache)
	if !global.TPA_CONFIG.System.UseMultipoint {
		cres.OkWithDetailed(wres.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.OpenID); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.OpenID); err != nil {
			global.TPA_LG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
			cres.FailWithMessage("设置登录状态失败", c)
			return
		}
		cres.OkWithDetailed(wres.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.TPA_LG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
		cres.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blackJWT); err != nil {
			cres.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.OpenID); err != nil {
			cres.FailWithMessage("设置登录状态失败", c)
			return
		}
		//设置用户缓存
		cres.OkWithDetailed(wres.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// HttpGetRequest HTTP GET请求带参数
func (u *WxLoginApi) HttpGetRequest(wx *request.WxCode) (wxLoginInfo wres.WXLoginResp, err error) {
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
