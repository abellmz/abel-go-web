package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

// 若申明为局部变量，就会只执行多次
var once sync.Once

// PrintOnce  这个方法不管调用几次都只执行一次
func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
