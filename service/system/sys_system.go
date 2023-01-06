package system

import (
	"go.uber.org/zap"
	"project/config"
	"project/global"
	"project/model/system"
	"project/utils"
)

type SystemConfigService struct{}

// GetSystemConfig @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetSystemConfig
// @description: 读取配置文件
// @return: err error, conf config.Server
func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.TPA_CONFIG
}

// SetSystemConfig @description   set system config,
// @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: SetSystemConfig
// @description: 设置配置文件
// @param: system model.System
// @return: err error
func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.TPA_VP.Set(k, v)
	}
	err = global.TPA_VP.WriteConfig()
	return err
}

// GetServerInfo @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetServerInfo
// @description: 获取服务器信息
// @return: server *utils.Server, err error
func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.TPA_LG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.TPA_LG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.TPA_LG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
