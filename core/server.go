package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/global"
	"project/initialize"
	"time"
)

type Server interface {
	ListenAndServe() error
}

// InitServer 初始化服务对象
func InitServer(address string, router *gin.Engine) Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunWindowsServer() {
	router := initialize.Routers()
	router.Static("/from-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.TPA_CONFIG.System.Addr)
	s := InitServer(address, router)

	time.Sleep(10 * time.Millisecond)
	global.TPA_LG.Info("server run success on", zap.String("address", address))
	global.TPA_LG.Error(s.ListenAndServe().Error())
}
