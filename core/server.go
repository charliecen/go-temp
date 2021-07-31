package core

import (
	"fmt"
	"go-temp/global"
	"go-temp/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 初始化路由
	Router := initialize.Routers()
	// 地址
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)

	global.LOG.Info("服务器启动成功：", zap.String("地址：", address))

	global.LOG.Error(s.ListenAndServe().Error())
}
