package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3, 4, 5}
	s1 := arr1[:]
	fmt.Printf("s1: %d, type: %T \n", s1, s1)

	// 切片定义
	var s2 []int
	// 切片是引用类型，默认值为nil
	fmt.Println(s2 == nil)
	// 分配内存空间
	s2 = make([]int, 3, 5) // 创建一个长度为3， 容量为5的空间
	s2[1] = 5
	fmt.Printf("s2: %d, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := []int{1, 2, 3, 4, 5}
	// 切片追加元素
	s3 = append(s3, 6, 7, 8)
	fmt.Printf("s2: %d, len: %d, cap: %d\n", s3, len(s3), cap(s3))
}
