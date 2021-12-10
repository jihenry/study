package main

import (
	"fmt"
	"regexp"
)

func main() {
	phone := "62081292619651"
	phoneMatchPattern := "^(?:62)?0?(\\d{10,14})$"
	rgx, err := regexp.Compile(phoneMatchPattern)
	if err != nil {
		panic(err)
	}
	ms := rgx.FindStringSubmatch(phone)
	amend := fmt.Sprintf("0062%s", ms[len(ms)-1])
	fmt.Printf("ms:%v amend:%s\n", ms, amend)
}
