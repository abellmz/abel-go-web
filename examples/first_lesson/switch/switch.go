package main

import "fmt"

func main() {
	ChooseFruit("草莓")
}
func ChooseFruit(fruit string) {
	switch fruit {
	case "苹果":
		fmt.Println("这是一个苹果")
	case "香蕉", "草莓":
		fmt.Println("这是香蕉或草莓")
	default:
		fmt.Printf("不知道是啥：%s \n", fruit)
	}
}
