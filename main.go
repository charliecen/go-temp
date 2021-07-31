package main

import (
	"go-temp/core"
	"go-temp/global"
	"go-temp/initialize"
)

func main() {
	// 初始化viper
	global.VP = core.Viper()
	// 初始化zap日志库
	global.LOG = core.Zap()
	// 初始化数据库
	global.DB = initialize.Gorm()
	// 初始化Redis
	initialize.Redis()
	// 初始化定时器
	initialize.Timer()
	// 启动服务
	core.RunServer()
}
