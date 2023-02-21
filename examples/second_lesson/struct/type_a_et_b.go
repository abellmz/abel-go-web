package main

import "fmt"

func main() {
	//  可以去除 "News" 结果一样
	var n News = fakeNews{
		Name: "hello",
	}
	n.Report()
}

type News struct {
	Name string
}

func (d News) Report() {
	fmt.Println("I am news: " + d.Name)
}

// type TypeA = TypeB,除了换个名字，没有任何区别
type fakeNews = News
