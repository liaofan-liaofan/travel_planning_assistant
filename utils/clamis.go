package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"project/global"
	"project/model/system"
	systemReq "project/model/system/request"
)

// GetRequestID 从Gin的Context中获取requestId
func GetRequestID(c *gin.Context) zap.Field {
	return zap.String("requestId", c.GetString("requestId"))
}

// GetUser 从Gin的Context中获取从jwt解析出来的用户信息
func GetUser(c *gin.Context) *systemReq.UserCache {
	if claims, exists := c.Get("claims"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的用户失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*systemReq.UserCache)
		return waitUse
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return uuid.UUID{}.String()
	} else {
		waitUse := claims.(*systemReq.UserCache)
		return waitUse.UUID
	}
}

// GetWxUserID 从Gin的Context中获取从jwt解析出来的用户ID(前台小程序用户的id)
func GetWxUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims-wx"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的小程序用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// GetWxUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID(前台小程序用户的id)
func GetWxUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims-wx"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的小程序用户UUID失败, 请检查路由是否使用jwt中间件!")
		return uuid.UUID{}.String()
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID.String()
	}
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.UserCache)
		return waitUse.ID
	}
}

// GetUserAuthority 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthority(c *gin.Context) []system.SysAuthority {
	if claims, exists := c.Get("claims"); !exists {
		global.TPA_LG.Error("从Gin的Context中获取从jwt解析出来的用户角色失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*systemReq.UserCache)
		return waitUse.Authority
	}
}
