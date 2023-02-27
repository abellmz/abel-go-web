package main

import "fmt"

func main() {

}

// Swimming 会游泳
type Swimming interface {
	Swim()
}

type Duck interface {
	// 鸭子会游泳，所以这里组合了它
	Swimming
}

type Base struct {
	Name string
}

type Concrete1 struct {
	Base
}

// 这种写法少见
type Concrete2 struct {
	*Base
}

func (c *Concrete1) SayHello() {
	//  c.Base.Name访问Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)
	//  也能直接访问Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)

	// 直接使用Base接收器的方法
	c.Base.SayHello()
	c.SayGoodBay()
}
func (b *Base) SayGoodBay() {

}

func (b *Base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}
