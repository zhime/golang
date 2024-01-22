package main

import "fmt"

func main() {
	// 指针定义
	var n *int
	num := 10
	n = &num
	fmt.Println(n, *n)
}
