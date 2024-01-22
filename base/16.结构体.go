package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	password string
}

type Account struct {
	User  User    `json:"user"`
	Money float32 `json:"money"`
	Name  string  `json:"name"`
}

func (u *User) GetName() {
	u.Name = "张三"
	fmt.Println("printName方法： ", u.Name)
	fmt.Printf("printName方法内部：%p \n", &u)
	//return u.Name
}

func main() {
	//var u1 = User{
	//	Name:     "zhangsan",
	//	Age:      18,
	//	password: "123456",
	//}
	//fmt.Printf("用户名:%s, 年龄:%d, 类型:%T\n", u1.Name, u1.Age, u1)
	var a1 = Account{
		User: User{
			Name:     "zhangsan",
			Age:      18,
			password: "123456",
		},
		Money: 19.99,
		Name:  "test",
	}
	a1.User.GetName()
	fmt.Println(a1.User.Name)

	u2 := User{
		Name: "lisi",
		Age:  20,
	}

	// 结构体转json
	jsonData, err := json.Marshal(u2)
	if err != nil {
		fmt.Println("结构体转json异常：", err)
	}
	fmt.Println(string(jsonData))

	// json转结构体
	var a2 Account
	jsonStr := `{"user":{"name":"hello","age":22},"money":19.88,"name":"dev"}`
	err = json.Unmarshal([]byte(jsonStr), &a2)
	if err != nil {
		fmt.Println("json转结构体异常：", err)
	}
	fmt.Println(a2)
}
