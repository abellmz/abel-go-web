package main

import "fmt"

func main() {
	name := "Tom"
	age := 17
	// 这个 API 是返回字符串的，所以大多数时候我们都是用这个
	str := fmt.Sprintf("hello, i am %s, i am %d years old \n", name, age)
	println(str)

	// 这个是直接输出，一般简单程序 DEBUG 会用它输出到一些信息到控制台
	fmt.Printf("hello, I am %s, I am %d years old \n", name, age)
	replaceHolder()
}
func replaceHolder() {
	u := &user{
		Name: "Tome",
		Age:  17,
	}
	// 按值的本来值输出
	fmt.Printf("v => %v \n", u)
	// 在%v的基础上，将结构体字段名和值展开
	fmt.Printf("+v => %+v \n", u)
	// 输出语法格式的值
	fmt.Printf("#v => %#v \n", u)
	// 输出语法格式的类型和值
	fmt.Printf("T => %T \n", u)
}

type user struct {
	Name string
	Age  int
}
