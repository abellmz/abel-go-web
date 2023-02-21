package main

import "fmt"

func main() {
	var p *ToyDuck = &ToyDuck{}
	var duck ToyDuck = *p
	duck.Swim()

	// 只是声明了，但是没有使用,会分配内存，但指针未赋值
	var nilDuck *ToyDuck
	if nilDuck == nil {
		fmt.Println("nilDuck is nil")
	}
}
