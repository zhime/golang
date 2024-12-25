package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/model"
	"net/http"
)

func getUser(c *gin.Context) {

	user := &model.User{
		Name:    "zhang",
		Age:     18,
		Phone:   132000000000,
		Address: "hangzhou",
		Email:   "admin@qq.com",
	}

	//jsonByte, err := json.Marshal(&user)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	fmt.Println(err)
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}

func addUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
