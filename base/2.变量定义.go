package main

import "fmt"

// 全局变量
var (
	v1 int
	v2 int = 2
	v3     = 3
)

func main() {
	// 只声明，不赋值，默认是0
	var num1 int
	fmt.Println(num1)

	// 先声明，再赋值
	var num2 int
	num2 = 10
	fmt.Println(num2)

	// 声明的同时赋值
	var num3 int = 11
	fmt.Println(num3)

	// 自动推断
	var num4 = 12
	fmt.Println(num4)

	// 声明赋值
	num5 := 13
	fmt.Println(num5)

	// 批量声明
	var (
		n1 int
		n2 int = 2
		n3     = 3
	)
	fmt.Println(n1, n2, n3)

	// 常量，声明之后不能修改
	const C int = 1
	fmt.Println(C)

	// iota
	const (
		a int = iota
		b     = 2
		c     = iota
		d     = iota
		e     = 1
		f     = iota
	)
	fmt.Println(a, b, c, d, e, f)
}
