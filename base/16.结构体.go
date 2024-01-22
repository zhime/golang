package main

import (
	"fmt"
)

type User struct {
	Name     string
	Age      int
	password string
}

func main() {
	var u1 = User{
		Name:     "zhangsan",
		Age:      18,
		password: "123456",
	}
	fmt.Printf("用户名:%s, 年龄:%d, 类型:%T\n", u1.Name, u1.Age, u1)
}
