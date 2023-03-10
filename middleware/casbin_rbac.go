package middleware

import (
	"github.com/gin-gonic/gin"
	"project/global"
	"project/model/common/response"
	"project/model/system/request"
	"strconv"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.UserCache)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 判断策略中是否存在
		success := false
		// 获取用户的角色
		for _, v := range waitUse.Authority {
			sub := strconv.Itoa(int(v.AuthorityId))
			success, _ = global.TPA_CASBIN.Enforce(sub, obj, act)
			if success {
				break
			}
		}
		if global.TPA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
