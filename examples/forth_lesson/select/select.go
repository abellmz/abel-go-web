package main

import (
	"fmt"
	"time"
)

func main() {
	// 这个不能在 main 函数运行，是因为运行起来，
	// 所有的goroutine都被我们搞sleep了，直接就崩了
	//Select()
}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg from channel1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "msg from channel2"
	}()

	//select 中case同时，顺序是没有保证的
	//map 中，key-value 也是没有顺序保证的
	// default 谨慎添加，因为很大可能一进入就default,底层io用的会比较多
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default:
			time.Sleep(time.Second)
		}

	}
}
