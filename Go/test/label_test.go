package main

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	a := 1
loop:
	for j := 1; j < 10; j++ {
		fmt.Printf("j=%d a=%d\n", j, a)
		switch j {
		case 3:
			fmt.Println("case continue label")
			continue loop
		case 5:
			fmt.Println("case break label")
			break loop
		case 4:
			fmt.Println("case break")
			break
		default:
			a++
		}
		fmt.Printf("break j=%d a=%d\n", j, a)
	}
}
