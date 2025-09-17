package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/middleware"
	"github.com/zhime/golang/model"
	"github.com/zhime/golang/service"
)

func getUser(c *gin.Context) {
	// 获取当前登录用户信息
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

func addUser(c *gin.Context) {
	// 检查管理员权限（简单实现）
	userEmail, exists := middleware.GetCurrentUserEmail(c)
	if !exists || userEmail != "admin@qq.com" {
		c.JSON(http.StatusForbidden, model.ErrorResponse{
			Code:    40301,
			Message: "需要管理员权限",
		})
		return
	}

	var registerReq model.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: "参数错误：" + err.Error(),
		})
		fmt.Println(err)
		return
	}

	// 使用认证服务注册新用户
	authService := service.GetAuthService()
	newUser, err := authService.Register(&registerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    40001,
			Message: err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model.APIResponse{
		Message: "用户创建成功",
		Data:    newUser.ToUserInfo(),
	})
}
