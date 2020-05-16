package main

import (
	"fmt"
	"reflect"
)

func printRT(rt interface{}) {
	v := reflect.ValueOf(rt).Elem()
	fmt.Println(v.Type())
	t := reflect.TypeOf(rt)
	fmt.Println(v, t.Name())
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		name := v.Type().Field(i).Name
		switch fv.Kind() {
		case reflect.Int32:
			fmt.Println(name, fv.Int())
		case reflect.String:
			fmt.Println(name, fv.String())
		case reflect.Slice:
			for i := 0; i < fv.Len(); i++ {
				fmt.Printf("%s[%d]=%t\n", name, i, fv.Index(i).Bool())
			}
		case reflect.Map:
			if fv.IsNil() {
				//fv.Type()是成员类型，fv.Type().Elem()是复合类型的子类型，取map得key类型使用Key()，
				//取struct成员的类型可以使用Type的Field方法，也可以使用遍历
				mapi := reflect.MakeMap(fv.Type())
				fv.Set(mapi)
				elem := reflect.New(fv.Type().Elem()).Elem()
				elem.SetInt(12)
				key := reflect.New(fv.Type().Key()).Elem()
				key.SetString("xxaaa")
				mapi.SetMapIndex(key, elem)
			}
		case reflect.Ptr:
			if fv.IsNil() {
				//1. new一个指针Value对象，这里其实包括两步，异步是创建一个具体类型对象，然后把对象的指针使用ValueOf来生成一个Value
				elem := reflect.New(fv.Type().Elem())
				fv.Set(elem)
				elem.Elem().SetInt(12)
				//2. 这里其实也是有两步生成，但是这次是直接那到具体类型对象的Value而不是指针，然后再取指针生成一个Value赋值给fv
				// elem := reflect.New(fv.Type().Elem()).Elem()
				// elem.SetInt(20)
				// fv.Set(elem.Addr())
			}
		}
	}
}

func printIsAddr() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)
	fmt.Println(a.CanAddr(), b.CanAddr(), c.CanAddr(), d.CanAddr())
}

func main() {
	type RT struct {
		a string
		b int32
		c []bool
		M map[string]int32
		P *int32
	}
	rti := &RT{c: []bool{true, false}}
	printRT(rti)
	printIsAddr()
	fmt.Printf("rtnew = %#v P:%d\n", *rti, *rti.P)
}
