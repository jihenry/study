package main

import (
	"fmt"
	"reflect"
)

type sp struct {
	p1 string
}

func (sp) pf(a int) string {
	return "pf"
}

func (sp) cf(a int) string {
	return "cf"
}

type sc struct {
	sp
	c1 int
}

type spi interface {
	pf(int) string
}

type sci interface {
	cf(int) string
}

func guestSType(v1 interface{}) {
	_, isSp := v1.(sp)
	fmt.Printf("v1 is sp:%v\n", isSp)
	_, isSc := v1.(sc)
	fmt.Printf("v1 is sc:%v\n", isSc)
}

func guestIType(v1 interface{}) {
	_, isSp := v1.(spi)
	fmt.Printf("v1 is spi:%v\n", isSp)
	_, isSc := v1.(sci)
	fmt.Printf("v1 is sci:%v\n", isSc)
}

func guestType(v1 interface{}) {
	switch v1.(type) {
	case sp:
		fmt.Println("guestType sp")
	case sc:
		fmt.Println("guestType sc")
	case sci:
		fmt.Println("guestType sci")
	case spi:
		fmt.Println("guestType spi")
	}
}

func main() {
	// e := sp{}
	// guestSType(e)
	// guestIType(e)
	// guestType(e)
	a := 1
	t1 := reflect.TypeOf(&a)

	type CSType struct {
		a int
		b string
		t reflect.Type
	}
	mc := make(map[CSType]bool)
	cd := CSType{a: 1, b: "222", t: t1}
	mc[cd] = true
	fmt.Printf("%#v\n", mc)
	cd2 := CSType{a: 1, b: "222", t: reflect.TypeOf(&a)}
	_, has := mc[cd2]
	fmt.Printf("has:%#v", has)
}