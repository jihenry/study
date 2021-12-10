package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	fmt.Println(string(b))
	fmt.Println(getGID())
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
