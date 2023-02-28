package main

import (
	"fmt"
	"time"
)

func main() {

}

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

// 断定MetricsFilterBuilder是否为FilterBuilder类型
// var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		// 在实现中主动调用下一个
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了%d 纳秒", end-start)
	}
}
