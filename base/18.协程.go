package main

import (
	"fmt"
	"sync"
)

// SendCode 发送验证码
func SendCode() {
	fmt.Println("发送验证码开始")
	//time.Sleep(3 * time.Second)
	fmt.Println("发送验证码完成！")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// 实现用户注册功能
	fmt.Println("用户注册校验完成")
	// 发送验证码
	//SendCode() // 会阻塞主线程
	wg.Add(1)
	go SendCode() // 会阻塞主线程
	fmt.Println("验证码已发送，请注意查收...")
	wg.Wait()
}
