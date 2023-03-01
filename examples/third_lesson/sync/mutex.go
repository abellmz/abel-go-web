package main

import (
	"fmt"
	"sync"
)

func main() {
	Mutex()
}

var mutex sync.Mutex
var rwMutex sync.RWMutex

// Mutex 锁，只有当确定没有任何读操作，再用这个
func Mutex() {
	mutex.Lock()
	// 解锁放在defer中，否则可能造成死锁
	defer mutex.Unlock()
	// my code
	fmt.Println("it's \"mutex\" code")
}

// RwMutex 读写锁  建议使用这种加锁方式
func RwMutex() {
	//加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 加写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()
}

// Failed1 不可重入：不可再次加锁
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 未解锁，又加锁，造成死锁。若只有一个goroutine，那么这会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

// Failed2 不可升级，拿了读锁后是不能变成写锁的
func Failed2() {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	//这一句会死锁,若只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}
