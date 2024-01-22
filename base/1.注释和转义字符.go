package main

import "fmt"

// Sum 函数注释
func Sum(num ...int) (sum int) {
	for _, value := range num {
		sum += value
	}
	return
}

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// 单行注释
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1))

	/*
		多行注释
	*/

	// 转义字符
	fmt.Println("\"hello world\"")
	fmt.Println(`\"hello world\"`)

	// 格式化字符串
	/*
		%v	按值的本来值输出
		%+v	在 %v 基础上，对结构体字段名和值进行展开
		%#v	输出 Go 语言语法格式的值
		%T	输出 Go 语言语法格式的类型和值
		%%	输出 % 本体
		%b（oOdxX）	整型的不同进制方式显示
		%U	Unicode 字符
		%s	字符串
		%d	整数
		%f	浮点数
		%p	指针，十六进制方式显示
	*/
	user := User{
		Name:    "张三",
		Age:     18,
		Address: "hangzhou",
	}
	fmt.Println(user)
	fmt.Printf("user 类型：%T\n", user)
}
