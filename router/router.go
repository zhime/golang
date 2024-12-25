package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/getUser", getUser)
		apiGroup.POST("/addUser")
	}
}
