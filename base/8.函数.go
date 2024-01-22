package main

import "fmt"

// 标准函数
func getSum(n1 int, n2 int) (sum int) {
	sum = n1 + n2
	return sum
}
func main() {
	res := getSum(1, 2)
	fmt.Println(res)

	// 匿名函数
	a, b := func(n1, n2 int) (sum, minus int) {
		sum = n1 + n2
		minus = n1 - n2
		return
	}(1, 2)
	fmt.Println(a, b)
}
