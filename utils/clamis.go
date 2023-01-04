package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetRequestID 从Gin的Context中获取requestId
func GetRequestID(c *gin.Context) zap.Field {
	return zap.String("requestId", c.GetString("requestId"))
}
