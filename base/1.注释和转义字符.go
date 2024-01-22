package main

import "fmt"

// Sum 函数注释
func Sum(num ...int) (sum int) {
	for _, value := range num {
		sum += value
	}
	return
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
}
