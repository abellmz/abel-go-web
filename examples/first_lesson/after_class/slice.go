package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)

}

func Add(s []int, index int, value int) []int {
	l := len(s)
	//result :=[]int{}
	//fmt.Println(result)
	if index < 0 {
	} else if index == 0 {
		tmp := append([]int{}, value)
		s = append(tmp, s[0:]...)
	} else if index <= l-1 && index > 0 {
		tmp := append([]int{}, s[index:]...)
		s = append(s[:index], value)
		s = append(s, tmp...)
	} else {
		s = append(s[:index], value)
	}
	fmt.Println(s)
	return s
}

func Delete(s []int, index int) []int {
	s = append(s[0:index], s[index+1:]...)
	fmt.Println(s)
	return s
}
