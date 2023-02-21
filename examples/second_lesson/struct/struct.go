package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck   是指针
	duck1 := &ToyDuck{}
	duck1.Swim()

	// 是结构体实例
	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck 也是指针 new->将内存分配好，并将其比特位置为0
	duck3 := new(ToyDuck)
	duck3.Swim()

	//---------以下是另一种写法，推荐上面的写法——直观

	// 当你声明这样的时候，Go 就帮你分配好内存，初始化好了,类似duck2   不用担心空指针的问题，因为它压根就不是指针
	var duck4 ToyDuck
	duck4.Swim()

	// duck5 就是一个指针了，分配好内存，但只能放指针，但指针无赋值
	//var duck5 *ToyDuck
	// 这边会直接panic 掉
	//duck5.Swim()

	// 赋值，初始化按字段名字赋值
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用
	duck7 := ToyDuck{"蓝色", 1024}
	duck7.Swim()

	// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "橘色"
	duck8.Swim()

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
	fmt.Printf("门前一条河，游过一群鸭，我是%s的，%d元一只\n", t.Color, t.Price)
}
