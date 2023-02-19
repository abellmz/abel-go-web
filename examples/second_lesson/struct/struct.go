package main

import "fmt"

func main() {

}

// Book 书籍  结构体定义
type Book struct {
	title, author string
	num, id       int
}

// ToyDuck 玩具鸭 结构体定义
type ToyDuck struct {
	Color string
	Price uint64
}

// 结构体接收器  结构体接收器内部永远不要修 改字段
func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s，%d一只\n", t.Color, t.Price)
}
