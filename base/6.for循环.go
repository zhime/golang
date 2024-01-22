package main

import "fmt"

func main() {
	// 标准for循环
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	// 无限循环
	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 2 {
			break // 跳出当前循环
		}
	}

	// 条件循环
	for i <= 5 {
		i++
		if i == 4 {
			continue // 结束本次循环
		}
		fmt.Println(i)
	}

	// 九九乘法表
	for i = 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}
