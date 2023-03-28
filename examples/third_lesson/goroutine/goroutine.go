package main

import (
	"fmt"
	"time"
)

func main() {
	GoRoutine()
}

func GoRoutine() {
	go func() {
		time.Sleep(10 * time.Second)
	}()
	fmt.Printf("I am here")
	fmt.Println("I am here")
}
