package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"project/global"
	"project/middleware"
	"project/model/common/response"
	"project/model/system"
	systemReq "project/model/system/request"
	systemRes "project/model/system/response"
	"project/utils"
	"time"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /api/base/login [post]

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)

	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if !store.Verify(l.CaptchaId, l.Captcha, true) {
	//	response.FailWithMessage("验证码错误", c)
	//	return
	//}
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	if err, user := userService.Login(u); err != nil {
		global.TPA_LG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.TPA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := systemReq.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.TPA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.TPA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "gsdPlus",                                             // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.TPA_LG.Error("获取token失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取token失败", c)
		return
	}
	var authorityIds []uint
	for _, authority := range user.Authorities {
		authorityIds = append(authorityIds, authority.AuthorityId)
	}
	userCache := systemReq.UserCache{
		ID:          user.ID,
		UUID:        user.UUID.String(),
		Authority:   user.Authorities,
		AuthorityId: authorityIds,
		DeptId:      user.DeptId,
	}
	_ = jwtService.SetRedisUserInfo(user.UUID.String(), userCache)
	if !global.TPA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.TPA_LG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.TPA_LG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		//设置用户缓存
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
