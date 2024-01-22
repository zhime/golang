package main

import "fmt"

func main() {
	// map定义
	var m1 map[string]string
	m1 = make(map[string]string, 2)
	m1["name"] = "zhang"
	m1["address"] = "hangzhou"
	fmt.Println(m1)

	// 简写
	m2 := map[string]string{
		"name":    "zhang",
		"address": "hangzhou",
	}
	fmt.Println(m2)

	// 查找key
	value, ok := m2["name"]
	if ok {
		// 如果存在，ok就为true
		fmt.Printf("value = %v\n", value)
	} else {
		fmt.Printf("key 不存在")
	}

	// 遍历map
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// 删除元素
	delete(m2, "name")
	fmt.Println(m2)

	// 切片中的元素为map类型
	var s1 = make([]map[string]string, 2)
	s1[0] = make(map[string]string)
	s1[0]["name"] = "zhang"
	s1[0]["address"] = "hangzhou"
	fmt.Println(s1)

	var s2 = make([]map[string]string, 2)
	s2[0] = map[string]string{
		"name":    "s2",
		"address": "0",
	}
	fmt.Println(s2, s2[0]["name"])

	// 简写
	s3 := []map[string]string{
		{
			"name":    "s3",
			"address": "0",
		},
		{
			"name":    "s3",
			"address": "4",
		},
	}
	fmt.Println(s3, s3[0]["address"])

	// map中的值为切片类型
	var m3 map[string][]string
	m3 = make(map[string][]string)
	m3["hobby"] = []string{"游泳", "跑步"}
	fmt.Println(m3)

	// 简写
	m4 := map[string][]string{
		"habby": {"羽毛球", "乒乓球"},
	}
	fmt.Println(m4)
}
