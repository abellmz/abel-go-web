package main

import "fmt"

func main() {
	i := 13
	a := func() {
		fmt.Println("i is %d \n", i)
	}
	a()

	//返回的是一个方法，需要再掉一次，才是真正执行了这个方法
	//fmt.Println(ReturnClosure("Tom")())
	fn := ReturnClosure("Tom")
	fmt.Println(fn())

	//延迟绑定
	Delay()
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		// 每次放入一个函数  引用传递，而不是值传递
		fns = append(fns, func() {
			fmt.Printf("Hello, this is: %d \n", i)
		})
	}

	// index,value  遍历中调用
	for _, fn := range fns {
		fn()
	}
}
