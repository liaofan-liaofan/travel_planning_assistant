package main

import (
	"travel_planning_assistant/core"
	"travel_planning_assistant/global"
	"travel_planning_assistant/initialize"
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
	global.TPA_VP = core.Viper()      // 初始化Viper
	global.TPA_LG = core.Zap()        // 初始化zap日志库
	global.TPA_DB = initialize.Gorm() // gorm连接数据库
}
