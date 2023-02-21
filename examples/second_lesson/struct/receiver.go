package main

import "fmt"

func main() {
	// 实例，因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	// 结构体接收器 不改变原值
	u.ChangeName("Tom Changed!")
	// 指针接收器 改变原值
	u.ChangeAge(18)
	fmt.Printf("%v \n", u)

	up := &User{
		Name: "Jerry",
		Age:  12,
	}
	// 因为 ChangeName 的接收器是结构体,所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed")
	up.ChangeAge(120)
	fmt.Printf("%v \n", up)

}

type User struct {
	Name string
	Age  int
}

// ChangeName 结构体接收器,不能直接调用该方法，需要先实例化，后调用
// 不会改变自身状态
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// ChangeAge 指针接收器  会改变自身状态
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

// Handle 若是 type A B，则以下那么写方法，因为没有指针的概念
type Handle func()

func (h Handle) hello() {

}
