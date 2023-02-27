package main

import "fmt"

func main() {
	son := Son{
		Parent{},
	}
	// java:I am Parent
	//golang: I am Son,但凡程序跑到了父级，后面调用的都是父级的方法
	son.SayHello()
}

type Parent struct {
}
type Son struct {
	Parent
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

func (p Parent) Name() string {
	return "Parent"
}

// 定义了自己的 Name() 方法
func (s Son) Name() string {
	return "Son"
}
