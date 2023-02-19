package main

import (
	"fmt"
)

type Book struct {
	title, author string
	num, id       int
}

func main() {
	// 变量交换
	//a := 1
	//b := 2
	//a ,b =b,a
	//fmt.Println(a)
	//fmt.Println(b)
	fmt.Println("Hello World!")

	//匿名变量不占空间，不分配内存
	//a, _ := ReturnData()
	//_, b := ReturnData()
	//fmt.Println(a, b)

	//  指针变量
	//num := 1
	//var p *int
	//p = &num
	//fmt.Println("num变量的地址为:", p)
	//fmt.Println("指针变量p的地址为:", &p)

	//xx :=make([]string,4,3)
	//fmt.Println(xx)

	var book1 = Book{}
	fmt.Println(book1)
}

func ReturnData() (int, int) {
	return 10, 20
}
