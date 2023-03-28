package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"project/config"
)

var (
	TPA_DB                  *gorm.DB
	TPA_REDIS               *redis.Client
	TPA_VP                  *viper.Viper
	TPA_LG                  *zap.Logger
	TPA_CONFIG              config.Server
	TPA_CASBIN              *casbin.SyncedEnforcer
	TPA_Concurrency_Control = &singleflight.Group{}
)
