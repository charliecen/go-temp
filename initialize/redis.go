package initialize

import (
	"context"
	"go-temp/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis ping 失败, 错误:", zap.Any("err", err))
	} else {
		global.LOG.Info("redis pong 响应:", zap.String("pong", pong))
		global.Redis = client
	}
}
