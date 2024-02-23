package sys

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

// GetHostIp 获取主机IP
func GetHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	return ip
}

// GetHostName 获取主机名
func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("get hostname err: ", err)
	}
	return hostname
}

// GetOS 获取操作系统
func GetOS() string {
	return runtime.GOOS
}
