package main

import "fmt"

type msgType uint8    // 自定义数据类型，混用时需进行类型转换
type msgType2 = uint8 // 类型别名, 混用时不需要类型转换
// byte是uint8的别名, rune是int32的别名

func main() {
	var n1 msgType
	var n2 uint8
	fmt.Printf("value: %v, type: %T\n", n1, n1)
	fmt.Printf("value: %v, type: %T\n", n2, n2)
	fmt.Printf("value: %v, type: %T\n", msgType(n2), msgType(n2))

	var n3 msgType2
	fmt.Printf("value: %v, type: %T\n", n3, n3)

	var n4 rune
	fmt.Printf("value: %v, type: %T\n", n4, n4)
}
