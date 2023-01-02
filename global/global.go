package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"travel_planning_assistant/config"
)

var (
	TPA_DB     *gorm.DB
	TPA_REDIS  *redis.Client
	TPA_VP     *viper.Viper
	TPA_LG     *zap.Logger
	TPA_CONFIG config.Server
)
