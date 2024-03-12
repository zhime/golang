package main

import (
	"fmt"
	"github.com/zhime/golang/web_app/logger"
	"github.com/zhime/golang/web_app/mysql"
	"github.com/zhime/golang/web_app/settings"
)

// Go Web开发通用脚手架

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init config err:", err)
	}
	// 2.初始化日志
	logger.Init()
	// 3.初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql err:", err)
	}
	// 4.初始化redis
	// 5.注册路由
	// 6.启动服务
}
