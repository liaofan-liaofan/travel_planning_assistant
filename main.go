package main

import (
	"database/sql"
	_ "net/http/pprof"
	"project/core"
	"project/global"
	"project/initialize"
)

//go:generate go mod tidy -go=1.19
//go:generate go mod download

// @title                       TravelPlanningAssistant API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.TPA_VP = core.Viper()            // 初始化Viper
	global.TPA_LG = core.Zap()              // 初始化zap日志库
	global.TPA_DB = initialize.Gorm()       // gorm连接数据库
	global.TPA_REDIS = initialize.Redis()   // 初始化redis
	global.TPA_CASBIN = initialize.Casbin() // 初始化casbin
	if global.TPA_DB != nil {
		//initialize.MysqlTables(global.TPA_DB) // 初始化自动迁移表
		//initialize.InitDB()                   // 初始化表数据
		// 程序结束前关闭数据链接
		db, _ := global.TPA_DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
	core.RunWindowsServer()
}
