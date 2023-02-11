package main

import "fmt"

func main() {
	var number = 10
	var count = 0
	for i := 0; i <= number; i++ {
		count = fibonacci(i)
		fmt.Println(count)
	}
}

// recursion
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
