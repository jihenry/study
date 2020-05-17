package main

import (
	"errors"
	"fmt"
	"unsafe"
)

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	vname1 int
	vname2 string
)

func forType() {
	var a = "菜鸟教程"
	var b = 10
	var d bool
	c := 12
	var g int
	fmt.Println(a, b, c, d)
	fmt.Println("Hello", g)
	println(a, b, c)
}

func forTest() {
	for true {
		println("无限循环")
	}
}

func constTest() {
	const (
		a        = iota
		b        = unsafe.Sizeof("ddd")
		c string = "hello"
		d        = len("ddddddd")
		e        = iota
	)
	println(a, b, c, d, e)
}

func arrayTest() {
	var a [10]int = [10]int{1, 2, 6, 7}
	println(a[2])
	var b = a[3]
	println(b)
}

func pointerTest() {
	var a = "333"
	println(&a)
	a = "ddd"
	println(a)
	var p *string
	p = &a
	println(*p)
	fmt.Printf("p指针地址: %x\n", p)
	var f *int
	println(f)
	if f == nil {
		println("nil")
	}
}

func max(num1 int, num2 int) (int, string) {
	if num1 > num2 {
		return num1, ""
	} else {
		return num2, ""
	}
}

func testFunc() {
	var a = 2
	var b = 3
	var ret, _ = max(a, b)
	println(ret)
}

func testStruct() {
	type Student struct {
		ddd string
		ccc int
	}
	var aaa Student
	aaa.ddd = "ddghddd"
	aaa.ccc = 1
	var llll *Student = &aaa
	fmt.Printf("ccc=%d ddd=%s", aaa.ccc, llll.ddd)
}

func testSlice() {
	// var d []int = make([]int,2)
	var d []int = []int{2, 3, 5}
	var dc [3]int = [3]int{2, 3, 5}
	fmt.Printf("type(d)=%T\n", d)
	fmt.Printf("type(dc)=%T\n", dc)
	println(cap(d))
	fmt.Println("d=", d) //使用println打印的是d的变量即tostring
	d = append(d, 88)    //必须有人接收，否则就是没有使用
	fmt.Println("d=", d)
}

func sliceTest() {
	arr := []int{6, 2, 3, 4, 5}
	s := arr[:]
	for e, v := range s {
		fmt.Println(s[e], v)
	}
	s1 := make([]int, 3)
	for e, v := range s1 {
		fmt.Println(s1[e], v)
	}
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func rangeTest() {
	nums := []int{2, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println("sum=", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	length := 0
	for range nums {
		length++
	}
	fmt.Println(length)
}

func mapTest() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	a := `xxxx\n`
	fmt.Println(a)

	/*使用键输出地图值 */
	for country, value := range countryCapitalMap {
		fmt.Println(country, "首都是", value)
	}

	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func recursionTest() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func fibonacciTest() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func convertTest() {
	var a int = 2
	var b int = 3
	var mean float32 = float32(a) / float32(b)
	fmt.Printf("mean的值为: %f\n", mean)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 0, nil
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	//多行注释
	strFormat := `Cannot proceed, the divider is zero.
	dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func errTest() {
	_, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

func main() {
	// errTest()
	// var a *int
	// fmt.Printf("a type=%T", a)
	// s := "123你好456" + "xxxxxxxxxxxxxxx"
	// a := []rune(s)
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }
	// for k, v := range a {
	// 	fmt.Println(k, v)
	// }
	// a := true
	// b := 1
	// if a && b == 1 {
	// 	fmt.Println("xxxxxxxxxx")
	// }
	sliceTest()
}
