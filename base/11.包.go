package main

import (
	f "fmt"                            // 包别名
	_ "github.com/go-sql-driver/mysql" // 引入某个包，但不直接使用包里的函数，而是调用该包里面的init函数
	// . 全部引入
	//. "github.com/go-sql-driver/mysql"
)

func main() {
	f.Println("Hello")
}
