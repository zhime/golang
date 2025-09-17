package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/middleware"
	"github.com/zhime/golang/model"
)

func InitRouter(r *gin.Engine) {
	// 添加中间件
	r.Use(middleware.CORSMiddleware())        // CORS支持
	r.Use(middleware.RequestLoggerMiddleware()) // 请求日志
	r.Use(gin.Recovery())                      // 恢复中间件

	// API路由组
	apiGroup := r.Group("/api")
	{
		// 初始化认证路由
		InitAuthRoutes(apiGroup)

		// 用户管理路由（需要认证）
		userGroup := apiGroup.Group("/", middleware.AuthMiddleware())
		{
			userGroup.GET("/getUser", getUser)
			userGroup.POST("/addUser", addUser)
		}

		// 公开路由（不需要认证）
		publicGroup := apiGroup.Group("/public")
		{
			publicGroup.GET("/health", healthCheck)
			publicGroup.GET("/version", getVersion)
		}
	}
}

// healthCheck 健康检查接口
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "OK",
		Data: gin.H{
			"status": "healthy",
			"service": "golang-auth-system",
		},
	})
}

// getVersion 版本信息接口
func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "OK",
		Data: gin.H{
			"version": "1.0.0",
			"build_time": "2024-01-01",
			"go_version": "1.19",
		},
	})
}
