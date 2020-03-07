package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

//SortFloat64FastV1 ...
func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	fmt.Printf("SortFloat64FastV1 b:%T len:%d cap:%d\n", b, len(b), cap(b))
	// 以int方式给float64排序
	sort.Ints(b)
}

//SortFloat64FastV2 ...
func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}

func main() {
	SortFloat64FastV1(a)
	fmt.Println(a)
}
