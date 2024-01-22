package main

import "fmt"

func main() {
	// 数组定义
	//var arr1 [3]int = [3]int{1, 2, 3}
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)
	fmt.Println(arr1[0])

	// 可以使用...自动推导数组的长度
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr2)

	// 二维数组
	var arr3 [2][3]int
	arr3[0][2] = 10
	fmt.Println(arr3)

	var arr4 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr4)
}
