package main

import "fmt"

func main() {
	Young(8)
	Young(100)

	IfUsingNewVariable(10, 200)
	IfUsingNewVariable(100, 30)
}

func Young(age int) {
	if age < 18 {
		fmt.Println("I am a child!")
	} else {
		fmt.Println("I am not a child")
	}
}

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("距离太远，不来了： %d\n", distance)
	} else {
		fmt.Printf("距离并不远，来一趟： %d\n", distance)
	}

	// 这里不能访问,变量distance作用域为结构体之内
	//fmt.Printf("距离是： %d\n", distance)
}
