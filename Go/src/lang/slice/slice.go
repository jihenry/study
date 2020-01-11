package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	for index := 0; index < v.NumMethod(); index++ {
		method := v.Method(index)
		fmt.Printf("method type:%s name:%s type2:%s", method.Type(), t.Method(index).Name, t.Method(index).Type)
	}
}

func main() {
	// s1 := []int{1, 2, 3}
	// s2 := s1
	// fmt.Println(s2, s1)
	// s2 = []int{2, 3, 4, 5}
	// fmt.Println(s2, s1)
	// var a1 = [2]int{2, 3}
	// a1[0] = 200
	// a2 := a1
	// a2[1] = 300
	// fmt.Println(a1, len(a1), a2)

	// s1 := "hello沙河"
	// count := 0
	// for i, c := range s1 {
	// 	fmt.Printf("%d:%q\n", i, c)
	// 	if unicode.Is(unicode.Han, c) {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	// var s struct {
	// 	a string
	// 	b string
	// }
	// s.a = "发发"
	// s.b = "第二个"
	// s1 := struct {
	// 	name string
	// 	age  string
	// }{name: "xxxx", age: "xxx"}
	// fmt.Printf("%#v %v", s, s1)

	// s1 := make([]int, 3, 5)
	// s1[4] = 100
	// fmt.Println(s1)

	// s2 := [10]int{1: 100, 9: 50}
	// fmt.Println(s2, len(s2))

	m1 := make(map[int]string, 20)
	m1[2] = "xxxx"
	fmt.Println(m1, len(m1))
	printMethod(student{})
}
