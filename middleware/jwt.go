package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/model/system"
	"project/model/system/request"
	"project/service"
	"project/utils"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var userWxService = service.ServiceGroupApp.WechatServiceGroup.UserLoginService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		//校验uuid是否有效
		err1, _ := userWxService.FindUserByUuid(claims.UUID.String())
		if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil && err1 != nil {
			_ = jwtService.JoinInBlacklist(system.JwtBlacklist{Jwt: token})
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		//token续期
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.TPA_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			//多点登录拦截
			if global.TPA_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.TPA_LG.Error("get redis jwt failed", zap.Any("err", err), utils.GetRequestID(c))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JoinInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		if c.Request.Header.Get("isWx") == "ok" {
			c.Set("claims-wx", claims)
			c.Next()
		}
		if userInfo, err := jwtService.GetRedisUserInfo(claims.UUID.String()); err != nil {
			//载入上下文
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
		} else {
			c.Set("claims", &userInfo)
			c.Set("tokenClaims", claims)
			c.Next()
		}
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.TPA_CONFIG.JWT.SigningKey),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.TPA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
