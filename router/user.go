package router

import (
	"encoding/json"
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

	jsonByte, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, string(jsonByte))
}
