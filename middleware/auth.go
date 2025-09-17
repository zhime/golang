package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/model"
	"github.com/zhime/golang/service"
	"github.com/zhime/golang/utils"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40103,
				Message: "用户未登录",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40102,
				Message: "Token格式错误",
			})
			c.Abort()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40102,
				Message: "Token不能为空",
			})
			c.Abort()
			return
		}

		// 检查token是否在黑名单中
		authService := service.GetAuthService()
		if authService.IsTokenBlacklisted(token) {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40102,
				Message: "Token已失效",
			})
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40102,
				Message: "Token无效或已过期",
			})
			c.Abort()
			return
		}

		// 验证用户是否存在
		user, err := authService.GetUserByID(claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Code:    40102,
				Message: "用户不存在",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_name", claims.Name)
		c.Set("user", user)
		c.Set("token", token)

		// 继续处理请求
		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件（不强制要求登录）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.Next()
			return
		}

		// 检查token是否在黑名单中
		authService := service.GetAuthService()
		if authService.IsTokenBlacklisted(token) {
			c.Next()
			return
		}

		// 尝试解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.Next()
			return
		}

		// 尝试获取用户信息
		user, err := authService.GetUserByID(claims.UserID)
		if err != nil {
			c.Next()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_name", claims.Name)
		c.Set("user", user)
		c.Set("token", token)

		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 首先必须通过认证
		AuthMiddleware()(c)
		if c.IsAborted() {
			return
		}

		// 获取用户信息
		userEmail, exists := c.Get("user_email")
		if !exists {
			c.JSON(http.StatusForbidden, model.ErrorResponse{
				Code:    40301,
				Message: "权限不足",
			})
			c.Abort()
			return
		}

		// 简单的管理员判断逻辑（实际项目中应该有更复杂的权限系统）
		email := userEmail.(string)
		if email != "admin@qq.com" {
			c.JSON(http.StatusForbidden, model.ErrorResponse{
				Code:    40301,
				Message: "需要管理员权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// RequestLoggerMiddleware 请求日志中间件
func RequestLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	})
}

// GetCurrentUser 从上下文中获取当前用户
func GetCurrentUser(c *gin.Context) (*model.User, bool) {
	if user, exists := c.Get("user"); exists {
		if u, ok := user.(*model.User); ok {
			return u, true
		}
	}
	return nil, false
}

// GetCurrentUserID 从上下文中获取当前用户ID
func GetCurrentUserID(c *gin.Context) (int, bool) {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(int); ok {
			return id, true
		}
	}
	return 0, false
}

// GetCurrentUserEmail 从上下文中获取当前用户邮箱
func GetCurrentUserEmail(c *gin.Context) (string, bool) {
	if userEmail, exists := c.Get("user_email"); exists {
		if email, ok := userEmail.(string); ok {
			return email, true
		}
	}
	return "", false
}