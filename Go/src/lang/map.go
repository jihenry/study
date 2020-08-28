package main

import "fmt"

func testRangeMap() {
	tmap := map[string]int64{
		"w":  14,
		"xx": 40,
	}
	for k, v := range tmap {
		fmt.Printf("k:%#v v:%#v\n", &k, &v)
	}
}
