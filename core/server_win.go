//go:build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 初始化服务对象
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
