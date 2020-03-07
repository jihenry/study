package main

import "fmt"

type spi interface {
	pf(int) string
}

type sci interface {
	cf(int) string
}

type sp struct {
	p1 string
}

type sc struct {
	c1 string
}

func (sc) cf(a int) string {
	return "cf"
}

func (sp) pf(a int) string {
	return "pf"
}

func (sp) cf(a int) string {
	return "cf"
}

func switchFunc(v interface{}) {
	switch v.(type) {
	case spi:
		fmt.Println("spi")
	case sci:
		fmt.Println("sci")
	default:
		fmt.Println("default")
	}
}

func main() {
	sp := sc{}
	switchFunc(sp)
}
