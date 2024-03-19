package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request.Method)
	//	fmt.Println(request.Host)
	//	fmt.Println(request.Header)
	//	_, _ = fmt.Fprintln(writer, "gin demo")
	//})
	//
	//err := http.ListenAndServe("0.0.0.0:9090", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := r.Group("api")
	{
		api.GET("user", Index)
	}
	_ = r.Run()
}
