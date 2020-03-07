package main

import "fmt"

type custInt int

func main() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	cintv := custInt(2)

	fmt.Printf("hello:%T, world:%T", hello, world)

	fmt.Printf("%#v\n", []rune("世界abb"))
	fmt.Printf("%#v\n", []rune(s))
	fmt.Printf("cintv:%T", cintv)
}
