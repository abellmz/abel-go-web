package main

import "fmt"

func main() {
	// 创建一个预估容量是3的 map
	m := make(map[string]string, 3)
	// 未指定预估容量
	m1 := make(map[string]string)
	// 直接初始化
	m2 := map[string]string{
		"Tom": "Jerry",
	}

	// 赋值
	m["hello"] = "world"
	m1["hello"] = "world"
	m2["hello"] = "world"
	val := m["hello"]
	println(val)

	// 使用两个返回值，第一个是key对应的value，第二个是是否存在这个key，值为true/false
	val, ok := m["hello"] // invalid_key
	if !ok {
		println("key not found")
	}
	println("ok-value: %v", ok)
	println(m["hello"])
	fmt.Println(m["hello"])
	for key, val := range m {
		fmt.Printf("%s => %s \n", key, val)
	}

	// 注意:当int作为value类型时，如果只用一个返回值，就无法判断到底是key对应的数量为0还是说不存在这个key
	fmt.Println(m)
	// 有符号整型int 默认int32,4字节，32位
	m3 := map[string]int{
		"watermelon": 0,
	}
	Fruits, isExist := m3["watermelon"]
	fmt.Println(Fruits, isExist)
	Fruits, isExist = m3["xxx"]
	fmt.Println(Fruits, isExist)

}
