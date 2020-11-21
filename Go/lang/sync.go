package main

import (
	"fmt"
	"sync"
)

func testSync() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
		// if err := recover(); err != nil {
		// 	fmt.Println("recover err:", err)
		// }
		// fmt.Println("3")
	}()
	var lock sync.Mutex
	lock.Lock()
	fmt.Println("main lock")
	lock.Unlock()
	lock.Unlock() //不能两次unlock
	fmt.Println("main unlock finish")
	// panic("this is a panic message")
}
