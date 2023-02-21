package main

import "fmt"

func main() {
	fake := FackFish{}
	fake.FakeSwim()
	// fake 编译错误,无法调用原来 Fish 的方法
	//fake.Swim()

	// 类型转换为Fish 有类似使用父方法，需要先装换一下
	td := Fish(fake)
	// 真的变成了鱼
	td.Swim()

	sFake := StrongFakeFish{}
	// 这里就是调用了自己的方法
	sFake.Swim()

	td = Fish(sFake)
	// 真的变成了鱼
	td.Swim()
}

type Fish struct{}

// FackFish 定义了一个新类型，注意是新类型   之前的类型是struct/interface,现在的类型是Fish
// 但是Fish来自于struct,所以新类型就是struct的延伸
type FackFish Fish

// FackFish 的结构体接收器
func (f FackFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎\n")
}

// 从Fish中 定义了一个新类型
type StrongFakeFish Fish

// StrongFakeFish 的结构体接收器
func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎\n")
}
func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是一只鸭子\n")
}
