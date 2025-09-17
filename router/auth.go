package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/middleware"
	"github.com/zhime/golang/model"
	"github.com/zhime/golang/service"
)

// AuthRouter 认证路由处理器
type AuthRouter struct {
	authService *service.AuthService
}

// NewAuthRouter 创建认证路由处理器
func NewAuthRouter() *AuthRouter {
	return &AuthRouter{
		authService: service.GetAuthService(),
	}
}

// InitAuthRoutes 初始化认证路由
func InitAuthRoutes(r *gin.RouterGroup) {
	authRouter := NewAuthRouter()
	
	// 认证相关路由
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authRouter.Login)        // 登录
		authGroup.POST("/register", authRouter.Register)  // 注册
		
		// 需要认证的路由
		authProtected := authGroup.Use(middleware.AuthMiddleware())
		{
			authProtected.POST("/logout", authRouter.Logout)                    // 注销
			authProtected.GET("/me", authRouter.GetCurrentUser)                 // 获取当前用户信息
			authProtected.PUT("/password", authRouter.ChangePassword)           // 修改密码
			authProtected.GET("/users", middleware.AdminMiddleware(), authRouter.GetAllUsers) // 获取所有用户（管理员）
		}
	}
}

// Login 用户登录
func (ar *AuthRouter) Login(c *gin.Context) {
	var loginReq model.LoginRequest
	
	// 绑定请求参数
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "参数错误：" + err.Error(),
		})
		return
	}

	// 执行登录
	loginResp, err := ar.authService.Login(&loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    40101,
			Message: err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "登录成功",
		Data:    loginResp,
	})
}

// Register 用户注册
func (ar *AuthRouter) Register(c *gin.Context) {
	var registerReq model.RegisterRequest
	
	// 绑定请求参数
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "参数错误：" + err.Error(),
		})
		return
	}

	// 执行注册
	user, err := ar.authService.Register(&registerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: err.Error(),
		})
		return
	}

	// 返回成功响应（不包含密码）
	c.JSON(http.StatusCreated, model.APIResponse{
		Message: "注册成功",
		Data:    user.ToUserInfo(),
	})
}

// Logout 用户注销
func (ar *AuthRouter) Logout(c *gin.Context) {
	// 从上下文获取token
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "Token不存在",
		})
		return
	}

	// 执行注销
	err := ar.authService.Logout(token.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "注销成功",
	})
}

// GetCurrentUser 获取当前用户信息
func (ar *AuthRouter) GetCurrentUser(c *gin.Context) {
	// 从中间件获取用户信息
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    40102,
			Message: "用户信息不存在",
		})
		return
	}

	// 返回用户信息（不包含密码）
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "获取成功",
		Data:    user.ToUserInfo(),
	})
}

// ChangePassword 修改密码
func (ar *AuthRouter) ChangePassword(c *gin.Context) {
	var changePasswordReq model.ChangePasswordRequest
	
	// 绑定请求参数
	if err := c.ShouldBindJSON(&changePasswordReq); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "参数错误：" + err.Error(),
		})
		return
	}

	// 获取当前用户ID
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    40102,
			Message: "用户信息不存在",
		})
		return
	}

	// 执行密码修改
	err := ar.authService.ChangePassword(userID, &changePasswordReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "密码修改成功",
	})
}

// GetAllUsers 获取所有用户（管理员功能）
func (ar *AuthRouter) GetAllUsers(c *gin.Context) {
	// 获取所有用户
	users := ar.authService.GetAllUsers()
	
	// 转换为UserInfo格式
	userInfos := make([]*model.UserInfo, len(users))
	for i, user := range users {
		userInfos[i] = user.ToUserInfo()
	}

	// 返回用户列表
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "获取成功",
		Data:    userInfos,
	})
}

// RefreshToken 刷新Token（可选功能）
func (ar *AuthRouter) RefreshToken(c *gin.Context) {
	// 从上下文获取token
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "Token不存在",
		})
		return
	}

	// 尝试刷新token
	newToken, err := ar.authService.RefreshToken(token.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: err.Error(),
		})
		return
	}

	// 返回新token
	c.JSON(http.StatusOK, model.APIResponse{
		Message: "刷新成功",
		Data: gin.H{
			"token": newToken,
		},
	})
}

// ValidateToken 验证Token有效性
func (ar *AuthRouter) ValidateToken(c *gin.Context) {
	// 如果能到达这里，说明token有效（通过了中间件验证）
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Code:    40102,
			Message: "Token无效",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Message: "Token有效",
		Data:    user.ToUserInfo(),
	})
}