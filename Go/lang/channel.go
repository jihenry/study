package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1(a chan int) {
	defer wg.Done()
	for index := 0; index < 100; index++ {
		a <- index
	}
	close(a)
}

func f2(a chan int, b chan int) {
	defer wg.Done()
	for x := range a {
		b <- x * x
	}
	close(b)
}

func testChannel() {
	wg.Add(2)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go f1(ch1)
	go f2(ch1, ch2)
	for x := range ch2 {
		fmt.Println(x)
	}
}
