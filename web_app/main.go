package main

import (
	"fmt"
	"github.com/zhime/golang/web_app/settings"
)

// Go Web开发通用脚手架

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("read config err:", err)
	}
	// 2.初始化日志
	// 3.初始化mysql
	// 4.初始化redis
	// 5.注册路由
	// 6.启动服务
}
