package main

import (
	"fmt"
	"sync"
)

func main() {
	// 并发时 线程安全map，但会降低性能
	m := sync.Map{}
	// 存入
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry1")

	// 取出
	val, ok := m.Load("mouse")
	if ok {
		// 断言
		fmt.Println(len(val.(string)))
	}
}
