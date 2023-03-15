package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &user{}
		}}
	// Get 返回的是 interface{}，所以需要类型断言
	u := pool.Get().(*user)
	// defer 还回去,它是一个临时对象存储池
	defer pool.Put(u)

	// 接着重置 u 这个对象
	u.Reset("Tom", "my_email@163.com")
	// 下边就是使用 u 来完成你的业务逻辑
	//fmt.Printf("hello")//yong用于debug
	fmt.Println(u) //用于控制台输出
}

type user struct {
	Name  string
	Email string
}

// Reset  一般来说，复用对象都要求我们取出来之后,重置里面的字段
func (u *user) Reset(name string, email string) {
	u.Email = email
	u.Name = name
}

func (u *user) P() {
	fmt.Println("1")
}
