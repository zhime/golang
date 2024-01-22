package main

import (
	"fmt"
)

func main() {
	// if else
	fmt.Println("请输入：")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("输入有误：", err)
	}
	if num < 60 {
		fmt.Println("不合格")
	} else if num < 80 {
		fmt.Println("良")
	} else {
		fmt.Println("优")
	}

	// switch case
	switch {
	case num < 60:
		fmt.Println("不合格")
	case num < 80:
		fmt.Println("良")
	default:
		fmt.Println("优")
	}

	switch num {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
		fallthrough // case 穿透
	default:
		fmt.Println("未知")
	}
}
