package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/golang/router"
)

func main() {
	//_ = cmd.Execute()

	r := gin.Default()
	router.InitRouter(r)
	_ = r.Run()
}
