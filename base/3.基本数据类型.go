package main

import "fmt"

func main() {
	// 整型，默认值0
	var num int
	fmt.Println(num)

	// 浮点型，默认值0
	var f float32
	fmt.Println(f)

	// 字符型, 默认值0
	var c byte
	c = 65
	fmt.Println(c)
	fmt.Printf("%s\n", string(c))

	// bool，默认false
	var bo bool
	fmt.Println(bo)

	// 字符串，默认空
	var s string
	s = "A"
	fmt.Println(s)
}
