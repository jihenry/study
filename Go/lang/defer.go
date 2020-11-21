package main

import "fmt"

func testDefer() {
	defer func() {
		fmt.Println("xxxxxxxxxxxxxxxxxxx")
		if err := recover(); err != nil {
			fmt.Println("xxx", err)
		}
	}()

	panic("PANIC")
}
