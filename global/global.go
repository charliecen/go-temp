package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go-temp/config"
	"go-temp/utils/timer"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	VP     *viper.Viper
	CONFIG config.Server
	LOG    *zap.Logger
	DB     *gorm.DB
	Timer  = timer.NewTimerTask()
	Redis  *redis.Client
	CC     = &singleflight.Group{}
)
