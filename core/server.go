package core

import (
	"fmt"
	"go.uber.org/zap"
	"project/global"
	"project/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	router := initialize.Routers()
	router.Static("/from-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.TPA_CONFIG.System.Addr)
	s := initServer(address, router)

	time.Sleep(10 * time.Millisecond)
	global.TPA_LG.Info("server run success on", zap.String("address", address))
	global.TPA_LG.Error(s.ListenAndServe().Error())
}
