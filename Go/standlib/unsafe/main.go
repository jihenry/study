package main

import (
	"fmt"
	"unsafe"
)

func float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }
func main() {
	fmt.Printf("%#x\n", float64bits(1.0)) // "0x3ff0000000000000"
	fmt.Printf("%#1x", 20)
}
