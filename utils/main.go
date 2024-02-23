package main

import (
	"fmt"
	"github.com/zhime/golang/utils/sys"
)

func main() {
	ip := sys.GetHostIp()
	fmt.Println(ip)

	hostname := sys.GetHostName()
	fmt.Println(hostname)

	os := sys.GetOS()
	fmt.Println(os)

	arch := sys.GetArch()
	fmt.Println(arch)
}
