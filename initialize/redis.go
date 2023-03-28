package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"os"
	"project/global"
)

func Redis() *redis.Client {
	redisCfg := global.TPA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.TPA_LG.Error("redis connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		global.TPA_LG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
